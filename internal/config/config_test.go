package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	// 创建临时配置文件
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")
	
	configContent := `
database:
  driver: "sqlite"
  dsn: "test.db"

output:
  dir: "./test_output"
  package: "testmodel"

tables:
  include: ["users", "products"]
  exclude: ["temp_*"]
  prefix: "t_"

options:
  generate_dao: true
  generate_sql: true
  json_tag: true
`
	
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}
	
	// 设置配置文件路径
	viper.Reset()
	viper.SetConfigFile(configFile)
	
	// 加载配置
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}
	
	// 验证配置
	if cfg.Database.Driver != "sqlite" {
		t.Errorf("期望数据库驱动为 'sqlite'，实际为 '%s'", cfg.Database.Driver)
	}
	
	if cfg.Database.DSN != "test.db" {
		t.Errorf("期望DSN为 'test.db'，实际为 '%s'", cfg.Database.DSN)
	}
	
	if cfg.Output.Dir != "./test_output" {
		t.Errorf("期望输出目录为 './test_output'，实际为 '%s'", cfg.Output.Dir)
	}
	
	if cfg.Output.Package != "testmodel" {
		t.Errorf("期望包名为 'testmodel'，实际为 '%s'", cfg.Output.Package)
	}
	
	if len(cfg.Tables.Include) != 2 {
		t.Errorf("期望包含2个表，实际为 %d", len(cfg.Tables.Include))
	}
	
	if !cfg.Options.GenerateDAO {
		t.Error("期望生成DAO为true")
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "有效配置",
			config: Config{
				Database: DatabaseConfig{
					Driver: "sqlite",
					DSN:    "test.db",
				},
				Output: OutputConfig{
					Dir:     "./output",
					Package: "model",
				},
			},
			wantErr: false,
		},
		{
			name: "缺少数据库驱动",
			config: Config{
				Database: DatabaseConfig{
					DSN: "test.db",
				},
				Output: OutputConfig{
					Dir:     "./output",
					Package: "model",
				},
			},
			wantErr: true,
			errMsg:  "数据库驱动不能为空",
		},
		{
			name: "缺少DSN",
			config: Config{
				Database: DatabaseConfig{
					Driver: "sqlite",
				},
				Output: OutputConfig{
					Dir:     "./output",
					Package: "model",
				},
			},
			wantErr: true,
			errMsg:  "数据库连接字符串不能为空",
		},
		{
			name: "不支持的数据库驱动",
			config: Config{
				Database: DatabaseConfig{
					Driver: "oracle",
					DSN:    "test.db",
				},
				Output: OutputConfig{
					Dir:     "./output",
					Package: "model",
				},
			},
			wantErr: true,
			errMsg:  "不支持的数据库驱动",
		},
		{
			name: "缺少输出目录",
			config: Config{
				Database: DatabaseConfig{
					Driver: "sqlite",
					DSN:    "test.db",
				},
				Output: OutputConfig{
					Package: "model",
				},
			},
			wantErr: true,
			errMsg:  "输出目录不能为空",
		},
		{
			name: "缺少包名",
			config: Config{
				Database: DatabaseConfig{
					Driver: "sqlite",
					DSN:    "test.db",
				},
				Output: OutputConfig{
					Dir: "./output",
				},
			},
			wantErr: true,
			errMsg:  "包名不能为空",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				if err == nil {
					t.Errorf("期望错误，但没有返回错误")
				} else if tt.errMsg != "" && err.Error() != tt.errMsg {
					t.Errorf("期望错误信息包含 '%s'，实际为 '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("不期望错误，但返回了错误: %v", err)
				}
			}
		})
	}
}

func TestSetDefaults(t *testing.T) {
	viper.Reset()
	setDefaults()
	
	if viper.GetString("output.dir") != "./generated" {
		t.Errorf("期望默认输出目录为 './generated'，实际为 '%s'", viper.GetString("output.dir"))
	}
	
	if viper.GetString("output.package") != "model" {
		t.Errorf("期望默认包名为 'model'，实际为 '%s'", viper.GetString("output.package"))
	}
	
	if !viper.GetBool("options.generate_dao") {
		t.Error("期望默认生成DAO为true")
	}
	
	if !viper.GetBool("options.generate_sql") {
		t.Error("期望默认生成SQL为true")
	}
	
	if !viper.GetBool("options.json_tag") {
		t.Error("期望默认生成JSON标签为true")
	}
}