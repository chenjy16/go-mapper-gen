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

// SQLGenerator SQL 生成器
type SQLGenerator struct {
	config *config.Config
}

// NewSQLGenerator 创建 SQL 生成器
func NewSQLGenerator(cfg *config.Config) *SQLGenerator {
	return &SQLGenerator{config: cfg}
}

// SQLData SQL 模板数据
type SQLData struct {
	TableName     string
	StructName    string
	Fields        []FieldData
	PrimaryKey    FieldData
	HasPrimaryKey bool
	InsertFields  []FieldData
	UpdateFields  []FieldData
}

// Generate 生成 SQL 代码
func (sg *SQLGenerator) Generate(table database.Table) error {
	// 准备模板数据
	data := sg.prepareTemplateData(table)
	
	// 生成代码
	code, err := sg.generateCode(data)
	if err != nil {
		return fmt.Errorf("生成代码失败: %w", err)
	}
	
	// 写入文件
	filename := fmt.Sprintf("%s.sql", toSnakeCase(data.StructName))
	filepath := filepath.Join(sg.config.Output.Dir, "sql", filename)
	
	if err := os.WriteFile(filepath, []byte(code), 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	
	fmt.Printf("  生成 SQL 文件: %s\n", filepath)
	return nil
}

// prepareTemplateData 准备模板数据
func (sg *SQLGenerator) prepareTemplateData(table database.Table) SQLData {
	structName := toPascalCase(removeTablePrefix(table.Name, sg.config.Tables.Prefix))
	
	data := SQLData{
		TableName:  table.Name,
		StructName: structName,
	}
	
	// 处理字段
	for _, col := range table.Columns {
		field := FieldData{
			Name:         col.Name,
			Type:         col.Type,
			Comment:      col.Comment,
			IsPrimaryKey: col.IsPrimaryKey,
		}
		
		if col.IsPrimaryKey {
			data.PrimaryKey = field
			data.HasPrimaryKey = true
		}
		
		// 非自增字段用于插入
		if !col.IsAutoIncr {
			data.InsertFields = append(data.InsertFields, field)
		}
		
		// 非主键字段用于更新
		if !col.IsPrimaryKey {
			data.UpdateFields = append(data.UpdateFields, field)
		}
		
		data.Fields = append(data.Fields, field)
	}
	
	return data
}

// generateCode 生成代码
func (sg *SQLGenerator) generateCode(data SQLData) (string, error) {
	tmpl := `-- {{ .StructName }} 相关 SQL 语句
-- 表名: {{ .TableName }}

-- 查询所有记录
SELECT {{ range $i, $field := .Fields }}{{ if $i }}, {{ end }}{{ $field.Name }}{{ end }}
FROM {{ .TableName }};

{{ if .HasPrimaryKey }}
-- 根据主键查询
SELECT {{ range $i, $field := .Fields }}{{ if $i }}, {{ end }}{{ $field.Name }}{{ end }}
FROM {{ .TableName }}
WHERE {{ .PrimaryKey.Name }} = ?;
{{ end }}

-- 插入记录
INSERT INTO {{ .TableName }} (
{{ range $i, $field := .InsertFields }}{{ if $i }},
{{ end }}    {{ $field.Name }}{{ end }}
) VALUES (
{{ range $i, $field := .InsertFields }}{{ if $i }},
{{ end }}    ?{{ end }}
);

-- 批量插入记录
INSERT INTO {{ .TableName }} (
{{ range $i, $field := .InsertFields }}{{ if $i }},
{{ end }}    {{ $field.Name }}{{ end }}
) VALUES {{ range $i := seq 3 }}{{ if $i }},{{ end }}
({{ range $j, $field := $.InsertFields }}{{ if $j }}, {{ end }}?{{ end }}){{ end }};

{{ if .HasPrimaryKey }}
-- 根据主键更新
UPDATE {{ .TableName }}
SET {{ range $i, $field := .UpdateFields }}{{ if $i }},
    {{ end }}{{ $field.Name }} = ?{{ end }}
WHERE {{ .PrimaryKey.Name }} = ?;

-- 根据主键删除
DELETE FROM {{ .TableName }}
WHERE {{ .PrimaryKey.Name }} = ?;

-- 批量删除
DELETE FROM {{ .TableName }}
WHERE {{ .PrimaryKey.Name }} IN (?, ?, ?);
{{ end }}

-- 分页查询
SELECT {{ range $i, $field := .Fields }}{{ if $i }}, {{ end }}{{ $field.Name }}{{ end }}
FROM {{ .TableName }}
ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.Name }}{{ else }}{{ (index .Fields 0).Name }}{{ end }}
LIMIT ? OFFSET ?;

-- 统计总数
SELECT COUNT(*) FROM {{ .TableName }};

{{ if .HasPrimaryKey }}
-- 检查记录是否存在
SELECT COUNT(*) FROM {{ .TableName }}
WHERE {{ .PrimaryKey.Name }} = ?;
{{ end }}

-- 条件查询示例
SELECT {{ range $i, $field := .Fields }}{{ if $i }}, {{ end }}{{ $field.Name }}{{ end }}
FROM {{ .TableName }}
WHERE 1=1
{{ range .Fields }}{{ if not .IsPrimaryKey }}  -- AND {{ .Name }} = ?
{{ end }}{{ end }}ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.Name }}{{ else }}{{ (index .Fields 0).Name }}{{ end }};

-- 模糊查询示例
SELECT {{ range $i, $field := .Fields }}{{ if $i }}, {{ end }}{{ $field.Name }}{{ end }}
FROM {{ .TableName }}
WHERE 1=1
{{ range .Fields }}{{ if or (contains .Type "varchar") (contains .Type "text") (contains .Type "char") }}  -- AND {{ .Name }} LIKE CONCAT('%', ?, '%')
{{ end }}{{ end }}ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.Name }}{{ else }}{{ (index .Fields 0).Name }}{{ end }};

-- 范围查询示例
SELECT {{ range $i, $field := .Fields }}{{ if $i }}, {{ end }}{{ $field.Name }}{{ end }}
FROM {{ .TableName }}
WHERE 1=1
{{ range .Fields }}{{ if or (contains .Type "int") (contains .Type "decimal") (contains .Type "float") (contains .Type "double") }}  -- AND {{ .Name }} BETWEEN ? AND ?
{{ end }}{{ end }}{{ range .Fields }}{{ if or (contains .Type "date") (contains .Type "time") }}  -- AND {{ .Name }} BETWEEN ? AND ?
{{ end }}{{ end }}ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.Name }}{{ else }}{{ (index .Fields 0).Name }}{{ end }};
`
	
	// 添加模板函数
	funcMap := template.FuncMap{
		"seq": func(n int) []int {
			result := make([]int, n)
			for i := range result {
				result[i] = i
			}
			return result
		},
		"contains": func(s, substr string) bool {
			return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
		},
	}
	
	t, err := template.New("sql").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("解析模板失败: %w", err)
	}
	
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("执行模板失败: %w", err)
	}
	
	return buf.String(), nil
}