package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Column 表示数据库列信息
type Column struct {
	Name         string `json:"name"`          // 列名
	Type         string `json:"type"`          // 数据库类型
	GoType       string `json:"go_type"`       // Go 类型
	Nullable     bool   `json:"nullable"`      // 是否可空
	IsPrimaryKey bool   `json:"is_primary_key"` // 是否主键
	IsAutoIncr   bool   `json:"is_auto_incr"`  // 是否自增
	DefaultValue string `json:"default_value"` // 默认值
	Comment      string `json:"comment"`       // 注释
}

// Table 表示数据库表信息
type Table struct {
	Name    string   `json:"name"`    // 表名
	Comment string   `json:"comment"` // 表注释
	Columns []Column `json:"columns"` // 列信息
}

// Database 数据库接口
type Database interface {
	Connect() error
	Close() error
	GetTables() ([]Table, error)
	GetTableColumns(tableName string) ([]Column, error)
}

// NewDatabase 创建数据库实例
func NewDatabase(driver, dsn string) (Database, error) {
	switch driver {
	case "mysql":
		return &MySQL{DSN: dsn}, nil
	case "postgres":
		return &PostgreSQL{DSN: dsn}, nil
	case "sqlite":
		return &SQLite{DSN: dsn}, nil
	default:
		return nil, fmt.Errorf("不支持的数据库驱动: %s", driver)
	}
}

// MySQL 实现
type MySQL struct {
	DSN string
	db  *sql.DB
}

func (m *MySQL) Connect() error {
	db, err := sql.Open("mysql", m.DSN)
	if err != nil {
		return fmt.Errorf("连接 MySQL 失败: %w", err)
	}
	
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping MySQL 失败: %w", err)
	}
	
	m.db = db
	return nil
}

func (m *MySQL) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}

func (m *MySQL) GetTables() ([]Table, error) {
	query := `
		SELECT 
			TABLE_NAME,
			TABLE_COMMENT
		FROM 
			INFORMATION_SCHEMA.TABLES 
		WHERE 
			TABLE_SCHEMA = DATABASE()
			AND TABLE_TYPE = 'BASE TABLE'
		ORDER BY 
			TABLE_NAME
	`
	
	rows, err := m.db.Query(query)
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
		columns, err := m.GetTableColumns(table.Name)
		if err != nil {
			return nil, fmt.Errorf("获取表 %s 的列信息失败: %w", table.Name, err)
		}
		table.Columns = columns
		
		tables = append(tables, table)
	}
	
	return tables, nil
}

func (m *MySQL) GetTableColumns(tableName string) ([]Column, error) {
	query := `
		SELECT 
			COLUMN_NAME,
			DATA_TYPE,
			IS_NULLABLE,
			COLUMN_KEY,
			EXTRA,
			COLUMN_DEFAULT,
			COLUMN_COMMENT
		FROM 
			INFORMATION_SCHEMA.COLUMNS 
		WHERE 
			TABLE_SCHEMA = DATABASE()
			AND TABLE_NAME = ?
		ORDER BY 
			ORDINAL_POSITION
	`
	
	rows, err := m.db.Query(query, tableName)
	if err != nil {
		return nil, fmt.Errorf("查询列信息失败: %w", err)
	}
	defer rows.Close()
	
	var columns []Column
	for rows.Next() {
		var col Column
		var nullable, columnKey, extra string
		var defaultValue sql.NullString
		
		if err := rows.Scan(
			&col.Name,
			&col.Type,
			&nullable,
			&columnKey,
			&extra,
			&defaultValue,
			&col.Comment,
		); err != nil {
			return nil, fmt.Errorf("扫描列信息失败: %w", err)
		}
		
		col.Nullable = nullable == "YES"
		col.IsPrimaryKey = columnKey == "PRI"
		col.IsAutoIncr = strings.Contains(extra, "auto_increment")
		if defaultValue.Valid {
			col.DefaultValue = defaultValue.String
		}
		
		// 转换为 Go 类型
		col.GoType = mysqlTypeToGoType(col.Type, col.Nullable)
		
		columns = append(columns, col)
	}
	
	return columns, nil
}

// mysqlTypeToGoType 将 MySQL 类型转换为 Go 类型
func mysqlTypeToGoType(mysqlType string, nullable bool) string {
	// 移除类型参数，如 varchar(255) -> varchar
	baseType := strings.Split(mysqlType, "(")[0]
	baseType = strings.ToLower(baseType)
	
	var goType string
	switch baseType {
	case "tinyint", "smallint", "mediumint", "int", "integer":
		goType = "int"
	case "bigint":
		goType = "int64"
	case "float":
		goType = "float32"
	case "double", "decimal", "numeric":
		goType = "float64"
	case "char", "varchar", "text", "tinytext", "mediumtext", "longtext":
		goType = "string"
	case "date", "datetime", "timestamp", "time":
		goType = "time.Time"
	case "tinyint(1)", "boolean", "bool":
		goType = "bool"
	case "json":
		goType = "json.RawMessage"
	case "blob", "tinyblob", "mediumblob", "longblob", "binary", "varbinary":
		goType = "[]byte"
	default:
		goType = "interface{}"
	}
	
	// 如果字段可空且不是指针类型，添加指针
	if nullable && !strings.HasPrefix(goType, "*") && goType != "interface{}" {
		goType = "*" + goType
	}
	
	return goType
}