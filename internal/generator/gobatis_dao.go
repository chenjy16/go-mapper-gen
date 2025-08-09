package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"go-mapper-gen/internal/config"
	"go-mapper-gen/internal/database"
)

// GobatisDAOGenerator gobatis DAO 生成器
type GobatisDAOGenerator struct {
	config *config.Config
}

// NewGobatisDAOGenerator 创建 gobatis DAO 生成器
func NewGobatisDAOGenerator(cfg *config.Config) *GobatisDAOGenerator {
	return &GobatisDAOGenerator{config: cfg}
}

// GobatisDAOData gobatis DAO 模板数据
type GobatisDAOData struct {
	Package       string
	ModelPackage  string
	DAOName       string
	StructName    string
	TableName     string
	PrimaryKey    FieldData
	Fields        []FieldData
	HasPrimaryKey bool
}

// Generate 生成 Gobatis DAO 代码
func (gdg *GobatisDAOGenerator) Generate(table database.Table, outputDir string) error {
	// 准备模板数据
	data := gdg.prepareTemplateData(table)
	
	// 生成接口代码
	interfaceCode, err := gdg.generateInterfaceCode(data)
	if err != nil {
		return fmt.Errorf("生成接口代码失败: %w", err)
	}
	
	// 确保输出目录存在
	daoDir := filepath.Join(outputDir, "dao")
	if err := os.MkdirAll(daoDir, 0755); err != nil {
		return fmt.Errorf("创建 DAO 目录失败: %w", err)
	}
	
	// 写入接口文件
	interfaceFile := filepath.Join(daoDir, fmt.Sprintf("%s_dao.go", toSnakeCase(data.StructName)))
	if err := os.WriteFile(interfaceFile, []byte(interfaceCode), 0644); err != nil {
		return fmt.Errorf("写入接口文件失败: %w", err)
	}
	
	return nil
}

// prepareTemplateData 准备模板数据
func (gdg *GobatisDAOGenerator) prepareTemplateData(table database.Table) GobatisDAOData {
	structName := toPascalCase(removeTablePrefix(table.Name, gdg.config.Tables.Prefix))
	
	data := GobatisDAOData{
		Package:      "dao",
		ModelPackage: "generated_gobatis/model",
		DAOName:      structName + "DAO",
		StructName:   structName,
		TableName:    table.Name,
	}
	
	// 处理字段
	for _, col := range table.Columns {
		field := FieldData{
			Name:         toPascalCase(col.Name),
			Type:         col.GoType,
			Comment:      col.Comment,
			IsPrimaryKey: col.IsPrimaryKey,
		}
		
		if col.IsPrimaryKey {
			data.PrimaryKey = field
			data.HasPrimaryKey = true
		}
		
		data.Fields = append(data.Fields, field)
	}
	
	return data
}

// generateInterfaceCode 生成接口代码
func (gdg *GobatisDAOGenerator) generateInterfaceCode(data GobatisDAOData) (string, error) {
	tmpl := `package {{ .Package }}

import (
	"context"
	model "{{ .ModelPackage }}"
)

// {{ .DAOName }} {{ .StructName }} 数据访问接口
// 符合 Gobatis 框架规范的方法命名和参数定义
type {{ .DAOName }} interface {
	// Insert 方法 - 插入操作
	// Insert 插入单个{{ .StructName }}记录
	Insert(ctx context.Context, record *model.{{ .StructName }}) error
	
	// InsertBatch 批量插入{{ .StructName }}记录
	InsertBatch(ctx context.Context, records []*model.{{ .StructName }}) error

{{ if .HasPrimaryKey }}
	// Select 方法 - 查询操作
	// SelectById 根据主键查询{{ .StructName }}
	SelectById(ctx context.Context, {{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (*model.{{ .StructName }}, error)
	
	// Update 方法 - 更新操作
	// UpdateById 根据主键更新{{ .StructName }}
	UpdateById(ctx context.Context, record *model.{{ .StructName }}) error
	
	// Delete 方法 - 删除操作
	// DeleteById 根据主键删除{{ .StructName }}
	DeleteById(ctx context.Context, {{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) error
	
	// DeleteByIds 根据主键列表批量删除{{ .StructName }}
	DeleteByIds(ctx context.Context, {{ toLower .PrimaryKey.Name }}s []{{ .PrimaryKey.Type }}) error
{{ end }}

	// SelectAll 查询所有{{ .StructName }}记录
	SelectAll(ctx context.Context) ([]*model.{{ .StructName }}, error)
	
	// SelectByPage 分页查询{{ .StructName }}记录
	SelectByPage(ctx context.Context, offset, limit int) ([]*model.{{ .StructName }}, error)
	
	// Count 统计{{ .StructName }}记录总数
	Count(ctx context.Context) (int64, error)
	
	// SelectByCondition 根据条件动态查询{{ .StructName }}记录
	SelectByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.{{ .StructName }}, error)
	
	// CountByCondition 根据条件统计{{ .StructName }}记录数
	CountByCondition(ctx context.Context, condition map[string]interface{}) (int64, error)

{{ if .HasPrimaryKey }}
	// ExistsById 检查指定主键的{{ .StructName }}记录是否存在
	ExistsById(ctx context.Context, {{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (bool, error)
{{ end }}

	// 兼容性方法 - 保持向后兼容
	// Create 创建{{ .StructName }} (兼容方法，内部调用 Insert)
	Create(ctx context.Context, {{ toLower .StructName }} *model.{{ .StructName }}) error

	// CreateBatch 批量创建{{ .StructName }} (兼容方法，内部调用 InsertBatch)
	CreateBatch(ctx context.Context, {{ toLower .StructName }}s []*model.{{ .StructName }}) error

{{ if .HasPrimaryKey }}
	// GetByID 根据ID获取{{ .StructName }} (兼容方法，内部调用 SelectById)
	GetByID(ctx context.Context, {{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (*model.{{ .StructName }}, error)

	// UpdateByID 根据ID更新{{ .StructName }} (兼容方法，内部调用 UpdateById)
	UpdateByID(ctx context.Context, {{ toLower .StructName }} *model.{{ .StructName }}) error

	// DeleteByID 根据ID删除{{ .StructName }} (兼容方法，内部调用 DeleteById)
	DeleteByID(ctx context.Context, {{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) error

	// DeleteByIDs 根据ID列表批量删除{{ .StructName }} (兼容方法，内部调用 DeleteByIds)
	DeleteByIDs(ctx context.Context, {{ toLower .PrimaryKey.Name }}s []{{ .PrimaryKey.Type }}) error

	// Exists 检查{{ .StructName }}是否存在 (兼容方法，内部调用 ExistsById)
	Exists(ctx context.Context, {{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (bool, error)
{{ end }}

	// GetAll 获取所有{{ .StructName }} (兼容方法，内部调用 SelectAll)
	GetAll(ctx context.Context) ([]*model.{{ .StructName }}, error)

	// GetByPage 分页获取{{ .StructName }} (兼容方法，内部调用 SelectByPage)
	GetByPage(ctx context.Context, offset, limit int) ([]*model.{{ .StructName }}, error)

	// FindByCondition 动态条件查询{{ .StructName }} (兼容方法，内部调用 SelectByCondition)
	FindByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.{{ .StructName }}, error)
}
`
	
	// 添加模板函数
	funcMap := template.FuncMap{
		"toLower":     strings.ToLower,
		"toSnakeCase": toSnakeCase,
	}
	
	t, err := template.New("gobatis_dao_interface").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("解析接口模板失败: %w", err)
	}
	
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("执行接口模板失败: %w", err)
	}
	
	return buf.String(), nil
}