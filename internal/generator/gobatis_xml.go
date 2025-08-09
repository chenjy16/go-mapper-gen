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

// GobatisXMLGenerator gobatis XML 映射文件生成器
type GobatisXMLGenerator struct {
	config *config.Config
}

// NewGobatisXMLGenerator 创建 gobatis XML 生成器
func NewGobatisXMLGenerator(cfg *config.Config) *GobatisXMLGenerator {
	return &GobatisXMLGenerator{config: cfg}
}

// GobatisXMLData gobatis XML 模板数据
type GobatisXMLData struct {
	Namespace       string
	DAOName         string
	StructName      string
	TableName       string
	PrimaryKey      FieldData
	Fields          []FieldData
	HasPrimaryKey   bool
	GenerateExample bool
}

// Generate 生成 gobatis XML 映射文件
func (gxg *GobatisXMLGenerator) Generate(table database.Table) error {
	// 准备模板数据
	data := gxg.prepareTemplateData(table)
	
	// 生成 XML 代码
	xmlCode, err := gxg.generateXMLCode(data)
	if err != nil {
		return fmt.Errorf("生成 XML 代码失败: %w", err)
	}
	
	// 写入 XML 文件
	filename := fmt.Sprintf("%s_mapper.xml", toSnakeCase(data.StructName))
	xmlPath := filepath.Join(gxg.config.Output.Dir, "mapper", filename)
	
	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(xmlPath), 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}
	
	if err := os.WriteFile(xmlPath, []byte(xmlCode), 0644); err != nil {
		return fmt.Errorf("写入 XML 文件失败: %w", err)
	}
	
	fmt.Printf("  生成 gobatis XML 映射文件: %s\n", xmlPath)
	return nil
}

// generateNamespace 生成namespace，支持自定义格式
func (gxg *GobatisXMLGenerator) generateNamespace(structName string) string {
	format := gxg.config.Options.NamespaceFormat
	if format == "" {
		format = "{struct}DAO" // 默认格式
	}
	
	// 替换占位符
	namespace := strings.ReplaceAll(format, "{struct}", structName)
	return namespace
}

// prepareTemplateData 准备模板数据
func (gxg *GobatisXMLGenerator) prepareTemplateData(table database.Table) GobatisXMLData {
	structName := toPascalCase(removeTablePrefix(table.Name, gxg.config.Tables.Prefix))
	
	// 生成namespace，支持自定义格式
	namespace := gxg.generateNamespace(structName)
	
	data := GobatisXMLData{
		Namespace:       namespace,
		DAOName:         structName + "DAO",
		StructName:      structName,
		TableName:       table.Name,
		GenerateExample: gxg.config.Options.GenerateExample,
	}
	
	// 处理字段
	for _, col := range table.Columns {
		field := FieldData{
			Name:         toPascalCase(col.Name),
			Type:         col.GoType,
			Comment:      col.Comment,
			IsPrimaryKey: col.IsPrimaryKey,
			ColumnName:   col.Name,
		}
		
		if col.IsPrimaryKey {
			data.PrimaryKey = field
			data.HasPrimaryKey = true
		}
		
		data.Fields = append(data.Fields, field)
	}
	
	return data
}

// generateXMLCode 生成 XML 代码
func (gxg *GobatisXMLGenerator) generateXMLCode(data GobatisXMLData) (string, error) {
	tmpl := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//gobatis.org//DTD Mapper 3.0//EN" "http://gobatis.org/dtd/gobatis-3-mapper.dtd">

<mapper namespace="{{ .Namespace }}">

    <!-- 结果映射 -->
    <resultMap id="{{ .StructName }}ResultMap" type="{{ .StructName }}">
        {{ range .Fields }}
        {{ if .IsPrimaryKey }}
        <id property="{{ .Name }}" column="{{ .ColumnName }}" />
        {{ else }}
        <result property="{{ .Name }}" column="{{ .ColumnName }}" />
        {{ end }}
        {{ end }}
    </resultMap>

    <!-- 基础字段列表 -->
    <sql id="Base_Column_List">
        {{ range $i, $field := .Fields }}{{ if $i }}, {{ end }}{{ $field.ColumnName }}{{ end }}
    </sql>

    <!-- 插入字段列表（不包含主键） -->
    <sql id="Insert_Column_List">
        {{- $first := true }}{{ range .Fields }}{{ if not .IsPrimaryKey }}{{ if not $first }}, {{ end }}{{ .ColumnName }}{{ $first = false }}{{ end }}{{ end }}
    </sql>

    <!-- 插入值列表（不包含主键） -->
    <sql id="Insert_Value_List">
        {{- $first := true }}{{ range .Fields }}{{ if not .IsPrimaryKey }}{{ if not $first }}, {{ end }}#{{"{"}}{{ .Name }}{{"}"}}{{ $first = false }}{{ end }}{{ end }}
    </sql>

    <!-- 更新字段列表（不包含主键） -->
    <sql id="Update_Set_List">
        {{- $first := true }}{{ range .Fields }}{{ if not .IsPrimaryKey }}{{ if not $first }}, {{ end }}{{ .ColumnName }} = #{{"{"}}{{ .Name }}{{"}"}}{{ $first = false }}{{ end }}{{ end }}
    </sql>

    <!-- Insert 方法 - 插入操作 -->
    <!-- Insert 插入单个{{ .StructName }}记录 -->
    <insert id="Insert" parameterType="{{ .StructName }}">
        INSERT INTO {{ .TableName }} (
            <include refid="Insert_Column_List" />
        ) VALUES (
            <include refid="Insert_Value_List" />
        )
    </insert>

    <!-- InsertBatch 批量插入{{ .StructName }}记录 -->
    <insert id="InsertBatch" parameterType="map">
        INSERT INTO {{ .TableName }} (
            <include refid="Insert_Column_List" />
        ) VALUES
        <foreach collection="records" item="item" separator=",">
            ({{- $first := true }}{{ range .Fields }}{{ if not .IsPrimaryKey }}{{ if not $first }}, {{ end }}#{{"{"}}item.{{ .Name }}{{"}"}}{{ $first = false }}{{ end }}{{ end }})
        </foreach>
    </insert>

    <!-- 兼容性方法 - Create -->
    <insert id="Create" parameterType="{{ .StructName }}">
        INSERT INTO {{ .TableName }} (
            <include refid="Insert_Column_List" />
        ) VALUES (
            <include refid="Insert_Value_List" />
        )
    </insert>

    <!-- 兼容性方法 - CreateBatch -->
    <insert id="CreateBatch" parameterType="map">
        INSERT INTO {{ .TableName }} (
            <include refid="Insert_Column_List" />
        ) VALUES
        <foreach collection="Items" item="item" separator=",">
            ({{- $first := true }}{{ range .Fields }}{{ if not .IsPrimaryKey }}{{ if not $first }}, {{ end }}#{{"{"}}item.{{ .Name }}{{"}"}}{{ $first = false }}{{ end }}{{ end }})
        </foreach>
    </insert>

    <!-- Select 方法 - 查询操作 -->
{{ if .HasPrimaryKey }}
    <!-- SelectById 根据ID查询{{ .StructName }}记录 -->
    <select id="SelectById" parameterType="{{ .PrimaryKey.Type }}" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </select>
{{ end }}

    <!-- SelectAll 查询所有{{ .StructName }}记录 -->
    <select id="SelectAll" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.ColumnName }}{{ else }}{{ (index .Fields 0).ColumnName }}{{ end }}
    </select>

    <!-- SelectByPage 分页查询{{ .StructName }}记录 -->
    <select id="SelectByPage" parameterType="map" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.ColumnName }}{{ else }}{{ (index .Fields 0).ColumnName }}{{ end }}
        LIMIT #{{"{"}}limit{{"}"}} OFFSET #{{"{"}}offset{{"}"}}
    </select>

    <!-- SelectByCondition 根据条件查询{{ .StructName }}记录 -->
    <select id="SelectByCondition" parameterType="map" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        <where>
            <if test="condition != null and condition != ''">
                ${condition}
            </if>
        </where>
        ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.ColumnName }}{{ else }}{{ (index .Fields 0).ColumnName }}{{ end }}
    </select>

    <!-- Update 方法 - 更新操作 -->
{{ if .HasPrimaryKey }}
    <!-- UpdateById 根据ID更新{{ .StructName }}记录 -->
    <update id="UpdateById" parameterType="{{ .StructName }}">
        UPDATE {{ .TableName }}
        SET <include refid="Update_Set_List" />
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </update>
{{ end }}

    <!-- Delete 方法 - 删除操作 -->
{{ if .HasPrimaryKey }}
    <!-- DeleteById 根据ID删除{{ .StructName }}记录 -->
    <delete id="DeleteById" parameterType="{{ .PrimaryKey.Type }}">
        DELETE FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </delete>

    <!-- DeleteByIds 根据ID列表批量删除{{ .StructName }}记录 -->
    <delete id="DeleteByIds" parameterType="map">
        DELETE FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} IN
        <foreach collection="ids" item="id" open="(" separator="," close=")">
            #{{"{"}}id{{"}"}}
        </foreach>
    </delete>

    <!-- ExistsById 检查指定ID的{{ .StructName }}记录是否存在 -->
    <select id="ExistsById" parameterType="{{ .PrimaryKey.Type }}" resultType="bool">
        SELECT COUNT(1) > 0
        FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </select>
{{ end }}

    <!-- Count 方法 - 计数操作 -->
    <!-- Count 获取{{ .StructName }}记录总数 -->
    <select id="Count" resultType="int64">
        SELECT COUNT(1)
        FROM {{ .TableName }}
    </select>

    <!-- CountByCondition 根据条件获取{{ .StructName }}记录数量 -->
    <select id="CountByCondition" parameterType="map" resultType="int64">
        SELECT COUNT(1)
        FROM {{ .TableName }}
        <where>
            <if test="condition != null and condition != ''">
                ${condition}
            </if>
        </where>
    </select>

    <!-- 兼容性方法 -->
{{ if .HasPrimaryKey }}
    <!-- 兼容性方法 - GetByID -->
    <select id="GetByID" parameterType="{{ .PrimaryKey.Type }}" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </select>

    <!-- 兼容性方法 - UpdateByID -->
    <update id="UpdateByID" parameterType="{{ .StructName }}">
        UPDATE {{ .TableName }}
        SET <include refid="Update_Set_List" />
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </update>

    <!-- 兼容性方法 - DeleteByID -->
    <delete id="DeleteByID" parameterType="{{ .PrimaryKey.Type }}">
        DELETE FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </delete>

    <!-- 兼容性方法 - DeleteByIDs -->
    <delete id="DeleteByIDs" parameterType="map">
        DELETE FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} IN
        <foreach collection="IDs" item="id" open="(" separator="," close=")">
            #{{"{"}}id{{"}"}}
        </foreach>
    </delete>

    <!-- 兼容性方法 - Exists -->
    <select id="Exists" parameterType="{{ .PrimaryKey.Type }}" resultType="int64">
        SELECT COUNT(1)
        FROM {{ .TableName }}
        WHERE {{ .PrimaryKey.ColumnName }} = #{{"{"}}{{ .PrimaryKey.Name }}{{"}"}}
    </select>
{{ end }}

    <!-- 兼容性方法 - GetAll -->
    <select id="GetAll" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.ColumnName }}{{ else }}{{ (index .Fields 0).ColumnName }}{{ end }}
    </select>

    <!-- 兼容性方法 - GetByPage -->
    <select id="GetByPage" parameterType="map" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.ColumnName }}{{ else }}{{ (index .Fields 0).ColumnName }}{{ end }}
        LIMIT #{{"{"}}limit{{"}"}} OFFSET #{{"{"}}offset{{"}"}}
    </select>

    <!-- 兼容性方法 - FindByCondition -->
    <select id="FindByCondition" parameterType="map" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        <where>
            {{ range .Fields }}
            <if test="{{ .Name }} != null{{ if eq .Type "string" }} and {{ .Name }} != ''{{ end }}">
                AND {{ .ColumnName }} = #{{"{"}}{{ .Name }}{{"}"}}
            </if>
            {{ end }}
        </where>
        ORDER BY {{ if .HasPrimaryKey }}{{ .PrimaryKey.ColumnName }}{{ else }}{{ (index .Fields 0).ColumnName }}{{ end }}
    </select>

{{ if .GenerateExample }}
    <!-- Example 方法 - 基于 Example 的查询操作 -->
    <!-- SelectByExample 根据Example条件查询{{ .StructName }}记录 -->
    <select id="SelectByExample" parameterType="github.com/chenjy16/gobatis/core/example.Example" resultMap="{{ .StructName }}ResultMap">
        SELECT 
            <include refid="Base_Column_List" />
        FROM {{ .TableName }}
        <where>
            <if test="criteria != null and criteria.size() > 0">
                <foreach collection="criteria" item="criterion" separator="AND">
                    <choose>
                        <when test="criterion.noValue">
                            ${criterion.condition}
                        </when>
                        <when test="criterion.singleValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}}
                        </when>
                        <when test="criterion.betweenValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}} AND #{{"{"}}criterion.secondValue{{"}"}}
                        </when>
                        <when test="criterion.listValue">
                            ${criterion.condition}
                            <foreach collection="criterion.value" item="listItem" open="(" separator="," close=")">
                                #{{"{"}}listItem{{"}"}}
                            </foreach>
                        </when>
                    </choose>
                </foreach>
            </if>
        </where>
        <if test="orderByClause != null and orderByClause != ''">
            ORDER BY ${orderByClause}
        </if>
        <if test="limit != null">
            LIMIT #{{"{"}}limit{{"}"}}
        </if>
        <if test="offset != null">
            OFFSET #{{"{"}}offset{{"}"}}
        </if>
    </select>

    <!-- CountByExample 根据Example条件统计{{ .StructName }}记录数 -->
    <select id="CountByExample" parameterType="github.com/chenjy16/gobatis/core/example.Example" resultType="int64">
        SELECT COUNT(1)
        FROM {{ .TableName }}
        <where>
            <if test="criteria != null and criteria.size() > 0">
                <foreach collection="criteria" item="criterion" separator="AND">
                    <choose>
                        <when test="criterion.noValue">
                            ${criterion.condition}
                        </when>
                        <when test="criterion.singleValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}}
                        </when>
                        <when test="criterion.betweenValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}} AND #{{"{"}}criterion.secondValue{{"}"}}
                        </when>
                        <when test="criterion.listValue">
                            ${criterion.condition}
                            <foreach collection="criterion.value" item="listItem" open="(" separator="," close=")">
                                #{{"{"}}listItem{{"}"}}
                            </foreach>
                        </when>
                    </choose>
                </foreach>
            </if>
        </where>
    </select>

    <!-- UpdateByExample 根据 Example 更新{{ .StructName }}记录 -->
    <update id="UpdateByExample" parameterType="map">
        UPDATE {{ .TableName }}
        <set>
            {{ range .Fields }}{{ if not .IsPrimaryKey }}
            <if test="record.{{ .Name }} != null">
                {{ .ColumnName }} = #{{"{"}}record.{{ .Name }}{{"}"}}{{ if not (eq . (index $.Fields (sub (len $.Fields) 1))) }},{{ end }}
            </if>
            {{ end }}{{ end }}
        </set>
        <where>
            <if test="example.criteria != null and example.criteria.size() > 0">
                <foreach collection="example.criteria" item="criterion" separator="AND">
                    <choose>
                        <when test="criterion.noValue">
                            ${criterion.condition}
                        </when>
                        <when test="criterion.singleValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}}
                        </when>
                        <when test="criterion.betweenValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}} AND #{{"{"}}criterion.secondValue{{"}"}}
                        </when>
                        <when test="criterion.listValue">
                            ${criterion.condition}
                            <foreach collection="criterion.value" item="listItem" open="(" separator="," close=")">
                                #{{"{"}}listItem{{"}"}}
                            </foreach>
                        </when>
                    </choose>
                </foreach>
            </if>
        </where>
    </update>

    <!-- DeleteByExample 根据Example条件删除{{ .StructName }}记录 -->
    <delete id="DeleteByExample" parameterType="github.com/chenjy16/gobatis/core/example.Example">
        DELETE FROM {{ .TableName }}
        <where>
            <if test="criteria != null and criteria.size() > 0">
                <foreach collection="criteria" item="criterion" separator="AND">
                    <choose>
                        <when test="criterion.noValue">
                            ${criterion.condition}
                        </when>
                        <when test="criterion.singleValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}}
                        </when>
                        <when test="criterion.betweenValue">
                            ${criterion.condition} #{{"{"}}criterion.value{{"}"}} AND #{{"{"}}criterion.secondValue{{"}"}}
                        </when>
                        <when test="criterion.listValue">
                            ${criterion.condition}
                            <foreach collection="criterion.value" item="listItem" open="(" separator="," close=")">
                                #{{"{"}}listItem{{"}"}}
                            </foreach>
                        </when>
                    </choose>
                </foreach>
            </if>
        </where>
    </delete>
{{ end }}

</mapper>
`
	
	// 添加模板函数
	funcMap := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}
	
	t, err := template.New("gobatis_xml").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("解析 XML 模板失败: %w", err)
	}
	
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("执行 XML 模板失败: %w", err)
	}
	
	return buf.String(), nil
}