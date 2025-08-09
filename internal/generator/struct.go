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

// StructGenerator 结构体生成器
type StructGenerator struct {
	config *config.Config
}

// NewStructGenerator 创建结构体生成器
func NewStructGenerator(cfg *config.Config) *StructGenerator {
	return &StructGenerator{config: cfg}
}

// StructData 结构体模板数据
type StructData struct {
	Package     string
	StructName  string
	TableName   string
	Comment     string
	Fields      []FieldData
	HasTimeType bool
	HasJSONType bool
}

// FieldData 字段模板数据
type FieldData struct {
	Name         string
	Type         string
	ColumnName   string
	JSONTag      string
	Comment      string
	IsPrimaryKey bool
}

// Generate 生成结构体代码
func (sg *StructGenerator) Generate(table database.Table) error {
	// 准备模板数据
	data := sg.prepareTemplateData(table)
	
	// 生成代码
	code, err := sg.generateCode(data)
	if err != nil {
		return fmt.Errorf("生成代码失败: %w", err)
	}
	
	// 写入文件
	filename := fmt.Sprintf("%s.go", toSnakeCase(data.StructName))
	modelDir := filepath.Join(sg.config.Output.Dir, "model")
	if err := os.MkdirAll(modelDir, 0755); err != nil {
		return fmt.Errorf("创建model目录失败: %w", err)
	}
	filepath := filepath.Join(modelDir, filename)
	
	if err := os.WriteFile(filepath, []byte(code), 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	
	fmt.Printf("  生成结构体文件: %s\n", filepath)
	return nil
}

// prepareTemplateData 准备模板数据
func (sg *StructGenerator) prepareTemplateData(table database.Table) StructData {
	data := StructData{
		Package:    "model",
		StructName: toPascalCase(removeTablePrefix(table.Name, sg.config.Tables.Prefix)),
		TableName:  table.Name,
		Comment:    table.Comment,
	}
	
	// 处理字段
	for _, col := range table.Columns {
		field := FieldData{
			Name:         toPascalCase(col.Name),
			Type:         col.GoType,
			ColumnName:   col.Name,
			Comment:      col.Comment,
			IsPrimaryKey: col.IsPrimaryKey,
		}
		
		// 生成标签
		var tags []string
		
		// 添加 db 标签（Gobatis 必需）
		tags = append(tags, fmt.Sprintf(`db:"%s"`, col.Name))
		
		// 生成 JSON 标签
		if sg.config.Options.JSONTag {
			tags = append(tags, fmt.Sprintf(`json:"%s"`, toSnakeCase(col.Name)))
		}
		
		if len(tags) > 0 {
			field.JSONTag = strings.Join(tags, " ")
		}
		

		
		// 检查是否需要导入特殊包
		if strings.Contains(col.GoType, "time.Time") {
			data.HasTimeType = true
		}
		if strings.Contains(col.GoType, "json.RawMessage") {
			data.HasJSONType = true
		}
		
		data.Fields = append(data.Fields, field)
	}
	
	return data
}



// generateCode 生成代码
func (sg *StructGenerator) generateCode(data StructData) (string, error) {
	tmpl := `package {{ .Package }}

{{ if or .HasTimeType .HasJSONType }}import ({{ if .HasTimeType }}
	"time"{{ end }}{{ if .HasJSONType }}
	"encoding/json"{{ end }}
){{ end }}

// {{ .StructName }} {{ .Comment }}
type {{ .StructName }} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }} ` + "`" + `{{ if .JSONTag }}{{ .JSONTag }}{{ end }}` + "`" + `{{ if .Comment }} // {{ .Comment }}{{ end }}
{{- end }}
}

// TableName 返回表名
func ({{ .StructName }}) TableName() string {
	return "{{ .TableName }}"
}
`
	
	t, err := template.New("struct").Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("解析模板失败: %w", err)
	}
	
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("执行模板失败: %w", err)
	}
	
	return buf.String(), nil
}

// 工具函数

// toPascalCase 转换为 PascalCase
func toPascalCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})
	
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	
	return strings.Join(words, "")
}

// toSnakeCase 转换为 snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

// removeTablePrefix 移除表前缀
func removeTablePrefix(tableName, prefix string) string {
	if prefix != "" && strings.HasPrefix(tableName, prefix) {
		return tableName[len(prefix):]
	}
	return tableName
}