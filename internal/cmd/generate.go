package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"go-mapper-gen/internal/config"
	"go-mapper-gen/internal/generator"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "生成 Go 代码",
	Long: `从数据库 schema 生成 Go 代码，包括：
- 结构体 (struct)
- DAO 层代码  
- SQL 语句
- CRUD 操作方法`,
	Run: func(cmd *cobra.Command, args []string) {
		runGenerate()
	},
}

func init() {
	// 数据库配置
	generateCmd.Flags().StringP("driver", "d", "", "数据库驱动 (mysql, postgres, sqlite)")
	generateCmd.Flags().String("dsn", "", "数据库连接字符串")
	
	// 输出配置
	generateCmd.Flags().StringP("output", "o", "./generated", "输出目录")
	generateCmd.Flags().StringP("package", "p", "model", "包名")
	
	// 表配置
	generateCmd.Flags().StringSlice("tables", []string{}, "要生成的表名 (逗号分隔)")
	generateCmd.Flags().StringSlice("exclude", []string{}, "要排除的表名 (逗号分隔)")
	generateCmd.Flags().String("prefix", "", "表前缀")
	
	// 生成选项
	generateCmd.Flags().Bool("dao", true, "生成 DAO 层代码")
	generateCmd.Flags().Bool("sql", true, "生成 SQL 语句")
	generateCmd.Flags().Bool("json-tag", true, "生成 JSON 标签")
	generateCmd.Flags().Bool("example", true, "生成 Example 方法 (支持 Gobatis v1.1.0)")
	
	// 绑定到 viper
	viper.BindPFlag("database.driver", generateCmd.Flags().Lookup("driver"))
	viper.BindPFlag("database.dsn", generateCmd.Flags().Lookup("dsn"))
	viper.BindPFlag("output.dir", generateCmd.Flags().Lookup("output"))
	viper.BindPFlag("output.package", generateCmd.Flags().Lookup("package"))
	viper.BindPFlag("tables.include", generateCmd.Flags().Lookup("tables"))
	viper.BindPFlag("tables.exclude", generateCmd.Flags().Lookup("exclude"))
	viper.BindPFlag("tables.prefix", generateCmd.Flags().Lookup("prefix"))
	viper.BindPFlag("options.generate_dao", generateCmd.Flags().Lookup("dao"))
	viper.BindPFlag("options.generate_sql", generateCmd.Flags().Lookup("sql"))
	viper.BindPFlag("options.json_tag", generateCmd.Flags().Lookup("json-tag"))
	viper.BindPFlag("options.generate_example", generateCmd.Flags().Lookup("example"))
}

func runGenerate() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	
	// 验证配置
	if err := cfg.Validate(); err != nil {
		log.Fatalf("配置验证失败: %v", err)
	}
	
	fmt.Printf("开始生成代码...\n")
	fmt.Printf("数据库: %s\n", cfg.Database.Driver)
	fmt.Printf("输出目录: %s\n", cfg.Output.Dir)
	fmt.Printf("包名: %s\n", cfg.Output.Package)
	
	// 创建生成器
	gen, err := generator.New(cfg)
	if err != nil {
		log.Fatalf("创建生成器失败: %v", err)
	}
	defer gen.Close()
	
	// 执行生成
	if err := gen.Generate(); err != nil {
		log.Fatalf("生成代码失败: %v", err)
	}
	
	fmt.Printf("代码生成完成！\n")
}