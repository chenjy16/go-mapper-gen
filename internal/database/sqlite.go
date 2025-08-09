package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// SQLite 实现
type SQLite struct {
	DSN string
	db  *sql.DB
}

func (s *SQLite) Connect() error {
	db, err := sql.Open("sqlite3", s.DSN)
	if err != nil {
		return fmt.Errorf("连接 SQLite 失败: %w", err)
	}
	
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping SQLite 失败: %w", err)
	}
	
	s.db = db
	return nil
}

func (s *SQLite) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *SQLite) GetTables() ([]Table, error) {
	query := `
		SELECT 
			name,
			'' as table_comment
		FROM 
			sqlite_master 
		WHERE 
			type = 'table'
			AND name NOT LIKE 'sqlite_%'
		ORDER BY 
			name
	`
	
	rows, err := s.db.Query(query)
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
		columns, err := s.GetTableColumns(table.Name)
		if err != nil {
			return nil, fmt.Errorf("获取表 %s 的列信息失败: %w", table.Name, err)
		}
		table.Columns = columns
		
		tables = append(tables, table)
	}
	
	return tables, nil
}

func (s *SQLite) GetTableColumns(tableName string) ([]Column, error) {
	// SQLite 使用 PRAGMA table_info 获取列信息
	query := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查询列信息失败: %w", err)
	}
	defer rows.Close()
	
	var columns []Column
	for rows.Next() {
		var col Column
		var cid int
		var notNull int
		var defaultValue sql.NullString
		
		if err := rows.Scan(
			&cid,
			&col.Name,
			&col.Type,
			&notNull,
			&defaultValue,
			&col.IsPrimaryKey,
		); err != nil {
			return nil, fmt.Errorf("扫描列信息失败: %w", err)
		}
		
		col.Nullable = notNull == 0
		if defaultValue.Valid {
			col.DefaultValue = defaultValue.String
		}
		
		// SQLite 自增检查需要额外查询
		col.IsAutoIncr = s.isAutoIncrement(tableName, col.Name)
		
		// 转换为 Go 类型
		col.GoType = sqliteTypeToGoType(col.Type, col.Nullable)
		
		columns = append(columns, col)
	}
	
	return columns, nil
}

// isAutoIncrement 检查列是否为自增
func (s *SQLite) isAutoIncrement(tableName, columnName string) bool {
	query := `
		SELECT sql 
		FROM sqlite_master 
		WHERE type = 'table' AND name = ?
	`
	
	var createSQL string
	err := s.db.QueryRow(query, tableName).Scan(&createSQL)
	if err != nil {
		return false
	}
	
	// 检查 CREATE TABLE 语句中该列是否包含 AUTOINCREMENT
	upperSQL := strings.ToUpper(createSQL)
	upperColumn := strings.ToUpper(columnName)
	
	// 查找列定义的位置
	columnIndex := strings.Index(upperSQL, upperColumn)
	if columnIndex == -1 {
		return false
	}
	
	// 查找该列定义的结束位置（下一个逗号或右括号）
	searchStart := columnIndex + len(upperColumn)
	nextComma := strings.Index(upperSQL[searchStart:], ",")
	nextParen := strings.Index(upperSQL[searchStart:], ")")
	
	endIndex := len(upperSQL)
	if nextComma != -1 && (nextParen == -1 || nextComma < nextParen) {
		endIndex = searchStart + nextComma
	} else if nextParen != -1 {
		endIndex = searchStart + nextParen
	}
	
	// 检查该列定义中是否包含 AUTOINCREMENT
	columnDef := upperSQL[columnIndex:endIndex]
	return strings.Contains(columnDef, "AUTOINCREMENT")
}

// sqliteTypeToGoType 将 SQLite 类型转换为 Go 类型
func sqliteTypeToGoType(sqliteType string, nullable bool) string {
	baseType := strings.ToLower(sqliteType)
	
	var goType string
	switch {
	case strings.Contains(baseType, "int"):
		if strings.Contains(baseType, "big") {
			goType = "int64"
		} else {
			goType = "int"
		}
	case strings.Contains(baseType, "real") || strings.Contains(baseType, "float") || strings.Contains(baseType, "double"):
		goType = "float64"
	case strings.Contains(baseType, "decimal") || strings.Contains(baseType, "numeric"):
		goType = "float64"
	case strings.Contains(baseType, "text") || strings.Contains(baseType, "char") || strings.Contains(baseType, "varchar"):
		goType = "string"
	case strings.Contains(baseType, "blob"):
		goType = "[]byte"
	case strings.Contains(baseType, "bool"):
		goType = "bool"
	case strings.Contains(baseType, "date") || strings.Contains(baseType, "time"):
		goType = "time.Time"
	default:
		// SQLite 的动态类型系统，默认使用 interface{}
		goType = "interface{}"
	}
	
	// 如果字段可空且不是指针类型，添加指针
	if nullable && !strings.HasPrefix(goType, "*") && goType != "interface{}" {
		goType = "*" + goType
	}
	
	return goType
}