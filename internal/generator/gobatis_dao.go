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
	Package         string
	ModelPackage    string
	DAOName         string
	StructName      string
	TableName       string
	PrimaryKey      FieldData
	Fields          []FieldData
	HasPrimaryKey   bool
	GenerateExample bool
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
	
	// 构建相对路径的 model 包导入路径
	modelPackage := "../model"
	
	data := GobatisDAOData{
		Package:         "dao",
		ModelPackage:    modelPackage,
		DAOName:         structName + "DAO",
		StructName:      structName,
		TableName:       table.Name,
		GenerateExample: gdg.config.Options.GenerateExample,
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
	model "{{ .ModelPackage }}"{{ if .GenerateExample }}
	"github.com/chenjy16/gobatis/core/example"{{ end }}
)

// {{ .DAOName }} {{ .StructName }} 数据访问接口
// 严格遵循 GoBatis 框架方法命名规则和返回值规范
type {{ .DAOName }} interface {
	// 插入方法 (INSERT) - 返回影响行数
	// Insert 插入单个{{ .StructName }}记录
	Insert(record *model.{{ .StructName }}) (int64, error)
	
	// InsertBatch 批量插入{{ .StructName }}记录
	InsertBatch(records []*model.{{ .StructName }}) (int64, error)
	
	// Add 添加{{ .StructName }}记录 (Insert 的别名)
	Add(record *model.{{ .StructName }}) (int64, error)
	
	// Create 创建{{ .StructName }}记录 (Insert 的别名)
	Create(record *model.{{ .StructName }}) (int64, error)
	
	// Save 保存{{ .StructName }}记录 (Insert 的别名)
	Save(record *model.{{ .StructName }}) (int64, error)

{{ if .HasPrimaryKey }}
	// 查询方法 (SELECT) - 返回查询结果
	// GetById 根据主键获取{{ .StructName }}
	GetById({{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (*model.{{ .StructName }}, error)
	
	// FindById 根据主键查找{{ .StructName }} (GetById 的别名)
	FindById({{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (*model.{{ .StructName }}, error)
	
	// SelectById 根据主键选择{{ .StructName }} (GetById 的别名)
	SelectById({{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (*model.{{ .StructName }}, error)
{{ end }}

	// GetAll 获取所有{{ .StructName }}记录
	GetAll() ([]*model.{{ .StructName }}, error)
	
	// FindAll 查找所有{{ .StructName }}记录 (GetAll 的别名)
	FindAll() ([]*model.{{ .StructName }}, error)
	
	// SelectAll 选择所有{{ .StructName }}记录 (GetAll 的别名)
	SelectAll() ([]*model.{{ .StructName }}, error)
	
	// ListAll 列出所有{{ .StructName }}记录 (GetAll 的别名)
	ListAll() ([]*model.{{ .StructName }}, error)
	
	// QueryAll 查询所有{{ .StructName }}记录 (GetAll 的别名)
	QueryAll() ([]*model.{{ .StructName }}, error)
	
	// GetByPage 分页获取{{ .StructName }}记录
	GetByPage(offset, limit int) ([]*model.{{ .StructName }}, error)
	
	// FindByPage 分页查找{{ .StructName }}记录 (GetByPage 的别名)
	FindByPage(offset, limit int) ([]*model.{{ .StructName }}, error)
	
	// SelectByPage 分页选择{{ .StructName }}记录 (GetByPage 的别名)
	SelectByPage(offset, limit int) ([]*model.{{ .StructName }}, error)
	
	// GetByCondition 根据条件获取{{ .StructName }}记录
	GetByCondition(condition map[string]interface{}) ([]*model.{{ .StructName }}, error)
	
	// FindByCondition 根据条件查找{{ .StructName }}记录 (GetByCondition 的别名)
	FindByCondition(condition map[string]interface{}) ([]*model.{{ .StructName }}, error)
	
	// SelectByCondition 根据条件选择{{ .StructName }}记录 (GetByCondition 的别名)
	SelectByCondition(condition map[string]interface{}) ([]*model.{{ .StructName }}, error)
	
	// QueryByCondition 根据条件查询{{ .StructName }}记录 (GetByCondition 的别名)
	QueryByCondition(condition map[string]interface{}) ([]*model.{{ .StructName }}, error)

	// 统计方法 - 返回数量
	// GetCount 获取{{ .StructName }}记录总数
	GetCount() (int64, error)
	
	// Count 统计{{ .StructName }}记录总数 (GetCount 的别名)
	Count() (int64, error)
	
	// CountByCondition 根据条件统计{{ .StructName }}记录数
	CountByCondition(condition map[string]interface{}) (int64, error)

{{ if .HasPrimaryKey }}
	// 存在性检查方法 - 返回布尔值
	// GetExistsById 检查指定主键的{{ .StructName }}记录是否存在
	GetExistsById({{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (bool, error)
{{ end }}

	// 更新方法 (UPDATE) - 返回影响行数
{{ if .HasPrimaryKey }}
	// UpdateById 根据主键更新{{ .StructName }}
	UpdateById(record *model.{{ .StructName }}) (int64, error)
	
	// ModifyById 根据主键修改{{ .StructName }} (UpdateById 的别名)
	ModifyById(record *model.{{ .StructName }}) (int64, error)
	
	// EditById 根据主键编辑{{ .StructName }} (UpdateById 的别名)
	EditById(record *model.{{ .StructName }}) (int64, error)
{{ end }}

	// UpdateByCondition 根据条件更新{{ .StructName }}记录
	UpdateByCondition(record *model.{{ .StructName }}, condition map[string]interface{}) (int64, error)

	// 删除方法 (DELETE) - 返回影响行数
{{ if .HasPrimaryKey }}
	// DeleteById 根据主键删除{{ .StructName }}
	DeleteById({{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (int64, error)
	
	// RemoveById 根据主键移除{{ .StructName }} (DeleteById 的别名)
	RemoveById({{ toLower .PrimaryKey.Name }} {{ .PrimaryKey.Type }}) (int64, error)
	
	// DeleteByIds 根据主键列表批量删除{{ .StructName }}
	DeleteByIds({{ toLower .PrimaryKey.Name }}s []{{ .PrimaryKey.Type }}) (int64, error)
	
	// RemoveByIds 根据主键列表批量移除{{ .StructName }} (DeleteByIds 的别名)
	RemoveByIds({{ toLower .PrimaryKey.Name }}s []{{ .PrimaryKey.Type }}) (int64, error)
{{ end }}

	// DeleteByCondition 根据条件删除{{ .StructName }}记录
	DeleteByCondition(condition map[string]interface{}) (int64, error)
	
	// RemoveByCondition 根据条件移除{{ .StructName }}记录 (DeleteByCondition 的别名)
	RemoveByCondition(condition map[string]interface{}) (int64, error)

{{ if .GenerateExample }}
	// Example 查询方法 - 支持 GoBatis Example 功能
	// GetByExample 根据 Example 条件获取{{ .StructName }}记录
	GetByExample(example *example.Example) ([]*model.{{ .StructName }}, error)
	
	// FindByExample 根据 Example 条件查找{{ .StructName }}记录 (GetByExample 的别名)
	FindByExample(example *example.Example) ([]*model.{{ .StructName }}, error)
	
	// SelectByExample 根据 Example 条件选择{{ .StructName }}记录 (GetByExample 的别名)
	SelectByExample(example *example.Example) ([]*model.{{ .StructName }}, error)
	
	// QueryByExample 根据 Example 条件查询{{ .StructName }}记录 (GetByExample 的别名)
	QueryByExample(example *example.Example) ([]*model.{{ .StructName }}, error)
	
	// ListByExample 根据 Example 条件列出{{ .StructName }}记录 (GetByExample 的别名)
	ListByExample(example *example.Example) ([]*model.{{ .StructName }}, error)
	
	// CountByExample 根据 Example 条件统计{{ .StructName }}记录数
	CountByExample(example *example.Example) (int64, error)
	
	// UpdateByExample 根据 Example 条件更新{{ .StructName }}记录
	UpdateByExample(record *model.{{ .StructName }}, example *example.Example) (int64, error)
	
	// ModifyByExample 根据 Example 条件修改{{ .StructName }}记录 (UpdateByExample 的别名)
	ModifyByExample(record *model.{{ .StructName }}, example *example.Example) (int64, error)
	
	// EditByExample 根据 Example 条件编辑{{ .StructName }}记录 (UpdateByExample 的别名)
	EditByExample(record *model.{{ .StructName }}, example *example.Example) (int64, error)
	
	// DeleteByExample 根据 Example 条件删除{{ .StructName }}记录
	DeleteByExample(example *example.Example) (int64, error)
	
	// RemoveByExample 根据 Example 条件移除{{ .StructName }}记录 (DeleteByExample 的别名)
	RemoveByExample(example *example.Example) (int64, error)
{{ end }}
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