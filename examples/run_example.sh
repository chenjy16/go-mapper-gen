#!/bin/bash

# Go Mapper Generator 示例运行脚本
echo "=== Go Mapper Generator 示例 ==="

# 检查是否存在可执行文件
if [ ! -f "../go-mapper-gen" ]; then
    echo "构建 go-mapper-gen..."
    cd ..
    go build -o go-mapper-gen ./cmd/go-mapper-gen
    cd examples
fi

# 创建示例数据库
echo "1. 创建示例数据库..."
sqlite3 example.db < init_database.sql
echo "✅ 数据库创建完成"

# 清理之前生成的代码
echo "2. 清理之前生成的代码..."
rm -rf generated/
echo "✅ 清理完成"

# 生成代码
echo "3. 生成代码..."
../go-mapper-gen generate --config example_config.yaml
echo "✅ 代码生成完成"

# 显示生成的文件结构
echo "4. 生成的文件结构:"
if [ -d "generated" ]; then
    tree generated/ 2>/dev/null || find generated/ -type f | sort
else
    echo "❌ 生成目录不存在"
fi

echo ""
echo "=== 示例完成 ==="
echo "生成的代码位于 examples/generated/ 目录下"
echo "数据库文件: examples/example.db"
echo ""
echo "你可以查看以下文件了解生成的代码:"
echo "- generated/model/*.go (数据模型)"
echo "- generated/dao/*.go (DAO 接口)"
echo "- generated/mapper/*.xml (XML 映射文件)"
echo "- generated/sql/*.sql (SQL 语句)"