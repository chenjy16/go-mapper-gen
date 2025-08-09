package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// PostgreSQL 实现
type PostgreSQL struct {
	DSN string
	db  *sql.DB
}

func (p *PostgreSQL) Connect() error {
	db, err := sql.Open("postgres", p.DSN)
	if err != nil {
		return fmt.Errorf("连接 PostgreSQL 失败: %w", err)
	}
	
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping PostgreSQL 失败: %w", err)
	}
	
	p.db = db
	return nil
}

func (p *PostgreSQL) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *PostgreSQL) GetTables() ([]Table, error) {
	query := `
		SELECT 
			table_name,
			COALESCE(obj_description(c.oid), '') as table_comment
		FROM 
			information_schema.tables t
		LEFT JOIN 
			pg_class c ON c.relname = t.table_name
		WHERE 
			t.table_schema = 'public'
			AND t.table_type = 'BASE TABLE'
		ORDER BY 
			t.table_name
	`
	
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查询表信息失败: %w", err)
	}
	defer rows.Close()
	
	var tables []Table
	for rows.Next() {
		var table Table
		if err := rows.Scan(&table.Name, &table.Comment); err != nil {
			return nil, fmt.Errorf("扫描表信息失败: %w", err)
		}
		
		// 获取列信息
		columns, err := p.GetTableColumns(table.Name)
		if err != nil {
			return nil, fmt.Errorf("获取表 %s 的列信息失败: %w", table.Name, err)
		}
		table.Columns = columns
		
		tables = append(tables, table)
	}
	
	return tables, nil
}

func (p *PostgreSQL) GetTableColumns(tableName string) ([]Column, error) {
	query := `
		SELECT 
			c.column_name,
			c.data_type,
			c.is_nullable,
			CASE WHEN pk.column_name IS NOT NULL THEN true ELSE false END as is_primary_key,
			CASE WHEN c.column_default LIKE 'nextval%' THEN true ELSE false END as is_auto_incr,
			COALESCE(c.column_default, '') as column_default,
			COALESCE(col_description(pgc.oid, c.ordinal_position), '') as column_comment
		FROM 
			information_schema.columns c
		LEFT JOIN 
			(
				SELECT ku.column_name
				FROM information_schema.table_constraints tc
				JOIN information_schema.key_column_usage ku
					ON tc.constraint_name = ku.constraint_name
				WHERE tc.constraint_type = 'PRIMARY KEY'
					AND tc.table_name = $1
			) pk ON c.column_name = pk.column_name
		LEFT JOIN 
			pg_class pgc ON pgc.relname = c.table_name
		WHERE 
			c.table_name = $1
			AND c.table_schema = 'public'
		ORDER BY 
			c.ordinal_position
	`
	
	rows, err := p.db.Query(query, tableName)
	if err != nil {
		return nil, fmt.Errorf("查询列信息失败: %w", err)
	}
	defer rows.Close()
	
	var columns []Column
	for rows.Next() {
		var col Column
		var nullable string
		
		if err := rows.Scan(
			&col.Name,
			&col.Type,
			&nullable,
			&col.IsPrimaryKey,
			&col.IsAutoIncr,
			&col.DefaultValue,
			&col.Comment,
		); err != nil {
			return nil, fmt.Errorf("扫描列信息失败: %w", err)
		}
		
		col.Nullable = nullable == "YES"
		
		// 转换为 Go 类型
		col.GoType = postgresTypeToGoType(col.Type, col.Nullable)
		
		columns = append(columns, col)
	}
	
	return columns, nil
}

// postgresTypeToGoType 将 PostgreSQL 类型转换为 Go 类型
func postgresTypeToGoType(pgType string, nullable bool) string {
	baseType := strings.ToLower(pgType)
	
	var goType string
	switch baseType {
	case "smallint", "integer", "int", "int4":
		goType = "int"
	case "bigint", "int8":
		goType = "int64"
	case "real", "float4":
		goType = "float32"
	case "double precision", "float8", "numeric", "decimal":
		goType = "float64"
	case "character varying", "varchar", "character", "char", "text":
		goType = "string"
	case "timestamp", "timestamp without time zone", "timestamp with time zone", "date", "time":
		goType = "time.Time"
	case "boolean", "bool":
		goType = "bool"
	case "json", "jsonb":
		goType = "json.RawMessage"
	case "bytea":
		goType = "[]byte"
	case "uuid":
		goType = "string"
	default:
		goType = "interface{}"
	}
	
	// 如果字段可空且不是指针类型，添加指针
	if nullable && !strings.HasPrefix(goType, "*") && goType != "interface{}" {
		goType = "*" + goType
	}
	
	return goType
}