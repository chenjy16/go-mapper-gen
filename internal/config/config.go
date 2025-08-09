package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config 生成器配置
type Config struct {
	Database DatabaseConfig `mapstructure:"database" yaml:"database"`
	Output   OutputConfig   `mapstructure:"output" yaml:"output"`
	Tables   TablesConfig   `mapstructure:"tables" yaml:"tables"`
	Options  OptionsConfig  `mapstructure:"options" yaml:"options"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver string `mapstructure:"driver" yaml:"driver"` // mysql, postgres, sqlite
	DSN    string `mapstructure:"dsn" yaml:"dsn"`       // 数据库连接字符串
}

// OutputConfig 输出配置
type OutputConfig struct {
	Dir     string `mapstructure:"dir" yaml:"dir"`         // 输出目录
	Package string `mapstructure:"package" yaml:"package"` // 包名
}

// TablesConfig 表配置
type TablesConfig struct {
	Include []string `mapstructure:"include" yaml:"include"` // 包含的表
	Exclude []string `mapstructure:"exclude" yaml:"exclude"` // 排除的表
	Prefix  string   `mapstructure:"prefix" yaml:"prefix"`   // 表前缀
}

// OptionsConfig 生成选项
type OptionsConfig struct {
	GenerateDAO     bool   `mapstructure:"generate_dao" yaml:"generate_dao"`         // 生成 DAO
	GenerateSQL     bool   `mapstructure:"generate_sql" yaml:"generate_sql"`         // 生成 SQL
	JSONTag         bool   `mapstructure:"json_tag" yaml:"json_tag"`                 // JSON 标签
	GenerateExample bool   `mapstructure:"generate_example" yaml:"generate_example"` // 生成 Example 方法
	NamespaceFormat string `mapstructure:"namespace_format" yaml:"namespace_format"` // XML namespace 格式模板，支持 {struct} 占位符
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	var cfg Config
	
	// 设置默认值
	setDefaults()
	
	// 解析配置
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	
	return &cfg, nil
}

// setDefaults 设置默认值
func setDefaults() {
	viper.SetDefault("output.dir", "./generated")
	viper.SetDefault("output.package", "model")
	viper.SetDefault("options.generate_dao", true)
	viper.SetDefault("options.generate_sql", true)
	viper.SetDefault("options.json_tag", true)
	viper.SetDefault("options.generate_example", true)
	viper.SetDefault("options.namespace_format", "{struct}DAO") // 默认格式：结构体名 + DAO
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Database.Driver == "" {
		return fmt.Errorf("数据库驱动不能为空")
	}
	
	if c.Database.DSN == "" {
		return fmt.Errorf("数据库连接字符串不能为空")
	}
	
	// 验证驱动类型
	supportedDrivers := []string{"mysql", "postgres", "sqlite"}
	if !contains(supportedDrivers, c.Database.Driver) {
		return fmt.Errorf("不支持的数据库驱动: %s, 支持的驱动: %s", 
			c.Database.Driver, strings.Join(supportedDrivers, ", "))
	}
	
	if c.Output.Dir == "" {
		return fmt.Errorf("输出目录不能为空")
	}
	
	if c.Output.Package == "" {
		return fmt.Errorf("包名不能为空")
	}
	
	return nil
}

// contains 检查切片是否包含指定元素
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}