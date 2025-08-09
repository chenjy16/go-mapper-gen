package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-mapper-gen/internal/config"
	"go-mapper-gen/internal/database"
)

// Generator 代码生成器
type Generator struct {
	config *config.Config
	db     database.Database
}

// New 创建新的生成器
func New(cfg *config.Config) (*Generator, error) {
	// 创建数据库连接
	db, err := database.NewDatabase(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("创建数据库连接失败: %w", err)
	}
	
	// 连接数据库
	if err := db.Connect(); err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}
	
	return &Generator{
		config: cfg,
		db:     db,
	}, nil
}

// Close 关闭生成器
func (g *Generator) Close() error {
	if g.db != nil {
		return g.db.Close()
	}
	return nil
}

// Generate 执行代码生成
func (g *Generator) Generate() error {
	// 获取所有表
	tables, err := g.db.GetTables()
	if err != nil {
		return fmt.Errorf("获取表信息失败: %w", err)
	}
	
	// 过滤表
	filteredTables := g.filterTables(tables)
	if len(filteredTables) == 0 {
		return fmt.Errorf("没有找到匹配的表")
	}
	
	fmt.Printf("找到 %d 个表需要生成代码\n", len(filteredTables))
	
	// 创建输出目录
	if err := g.createOutputDirs(); err != nil {
		return fmt.Errorf("创建输出目录失败: %w", err)
	}
	
	// 生成代码
	for _, table := range filteredTables {
		fmt.Printf("正在生成表 %s 的代码...\n", table.Name)
		
		// 生成结构体
		if err := g.generateStruct(table); err != nil {
			return fmt.Errorf("生成表 %s 的结构体失败: %w", table.Name, err)
		}
		
		// 生成 DAO 和 XML 映射文件
		if g.config.Options.GenerateDAO {
			if err := g.generateDAO(table); err != nil {
				return fmt.Errorf("生成表 %s 的 DAO 失败: %w", table.Name, err)
			}
		}
		
		// 生成 SQL
		if g.config.Options.GenerateSQL {
			if err := g.generateSQL(table); err != nil {
				return fmt.Errorf("生成表 %s 的 SQL 失败: %w", table.Name, err)
			}
		}
	}
	
	return nil
}

// filterTables 过滤表
func (g *Generator) filterTables(tables []database.Table) []database.Table {
	var filtered []database.Table
	
	for _, table := range tables {
		// 检查包含列表
		if len(g.config.Tables.Include) > 0 {
			found := false
			for _, include := range g.config.Tables.Include {
				if table.Name == include {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		
		// 检查排除列表
		excluded := false
		for _, exclude := range g.config.Tables.Exclude {
			if strings.Contains(exclude, "*") {
				// 支持通配符
				pattern := strings.ReplaceAll(exclude, "*", "")
				if strings.Contains(table.Name, pattern) {
					excluded = true
					break
				}
			} else if table.Name == exclude {
				excluded = true
				break
			}
		}
		if excluded {
			continue
		}
		
		// 检查前缀
		if g.config.Tables.Prefix != "" {
			if !strings.HasPrefix(table.Name, g.config.Tables.Prefix) {
				continue
			}
		}
		
		filtered = append(filtered, table)
	}
	
	return filtered
}

// createOutputDirs 创建输出目录
func (g *Generator) createOutputDirs() error {
	dirs := []string{
		g.config.Output.Dir,
	}
	
	if g.config.Options.GenerateDAO {
		dirs = append(dirs, filepath.Join(g.config.Output.Dir, "dao"))
		dirs = append(dirs, filepath.Join(g.config.Output.Dir, "mapper"))
	}
	
	if g.config.Options.GenerateSQL {
		dirs = append(dirs, filepath.Join(g.config.Output.Dir, "sql"))
	}
	
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("创建目录 %s 失败: %w", dir, err)
		}
	}
	
	return nil
}

// generateStruct 生成结构体
func (g *Generator) generateStruct(table database.Table) error {
	structGen := NewStructGenerator(g.config)
	return structGen.Generate(table)
}

// generateDAO 生成 DAO
func (g *Generator) generateDAO(table database.Table) error {
	// 生成 gobatis DAO 接口
	gobatisDAOGen := NewGobatisDAOGenerator(g.config)
	if err := gobatisDAOGen.Generate(table, g.config.Output.Dir); err != nil {
		return err
	}
	
	// 生成 gobatis XML 映射文件
	gobatisXMLGen := NewGobatisXMLGenerator(g.config)
	return gobatisXMLGen.Generate(table)
}

// generateGobatisDAO 生成 Gobatis DAO
func (g *Generator) generateGobatisDAO(cfg *config.Config, tables []database.Table) error {
	for _, table := range tables {
		gobatisDAOGen := NewGobatisDAOGenerator(cfg)
		if err := gobatisDAOGen.Generate(table, cfg.Output.Dir); err != nil {
			return fmt.Errorf("生成表 %s 的 Gobatis DAO 失败: %w", table.Name, err)
		}
		
		gobatisXMLGen := NewGobatisXMLGenerator(cfg)
		if err := gobatisXMLGen.Generate(table); err != nil {
			return fmt.Errorf("生成表 %s 的 Gobatis XML 失败: %w", table.Name, err)
		}
	}
	return nil
}

// generateSQL 生成 SQL
func (g *Generator) generateSQL(table database.Table) error {
	sqlGen := NewSQLGenerator(g.config)
	return sqlGen.Generate(table)
}