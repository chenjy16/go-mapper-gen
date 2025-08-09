# Go Mapper Generator

一个强大的 Go 代码生成工具，可以从数据库表结构自动生成 Go 结构体、DAO 层代码和 SQL 语句。支持 MySQL和 SQLite 数据库，专为 Gobatis 框架设计。

## 功能特性

- 🚀 **多数据库支持**: MySQL、PostgreSQL、SQLite
- 📝 **代码生成**: 自动生成 Go 结构体、DAO 接口和实现
- 🏷️ **标签支持**: JSON 标签自动生成
- 🔧 **高度可配置**: 支持表过滤、字段映射、命名规则等
- 📊 **SQL 生成**: 自动生成常用的 CRUD SQL 语句
- 🎯 **Gobatis 支持**: 专为 Gobatis 框架设计，生成完整的 DAO 接口、实现和 XML 映射文件

## 安装

从源码构建：

```bash
git clone https://github.com/your-username/go-mapper-gen.git
cd go-mapper-gen
go build -o go-mapper-gen ./cmd/go-mapper-gen
```

或者直接运行：

```bash
go run ./cmd/go-mapper-gen
```

## 使用方法

### 1. 配置文件方式

创建配置文件 `generator.yaml`:

```yaml
database:
  driver: mysql
  dsn: "user:password@tcp(localhost:3306)/database"
  
output:
  dir: "./generated"
  package: "model"
  
tables:
  include: ["users", "orders"]
  exclude: ["temp_*"]
  prefix: "t_"
  
options:
  generate_dao: true    # 生成 DAO 层代码
  generate_sql: true    # 生成 SQL 语句
  json_tag: true        # 生成 JSON 标签
```

然后运行：

```bash
go-mapper-gen generate -c generator.yaml
```

### 2. 命令行方式

```bash
go-mapper-gen generate \
  --driver mysql \
  --dsn "user:password@tcp(localhost:3306)/database" \
  --output ./generated \
  --package model \
  --tables users,orders
```

### 3. go:generate 集成

在你的 Go 文件中添加：

```go
//go:generate go-mapper-gen generate -c generator.yaml
```

然后运行：

```bash
go generate
```

### 命令行使用

```bash
# 使用配置文件生成代码
./go-mapper-gen generate -c generator-sqlite.yaml

# 直接指定参数生成代码
./go-mapper-gen generate --driver sqlite --dsn test.db --output-dir generated --package model

# 查看帮助
./go-mapper-gen --help
```

### 命令行参数

```bash
# 基本用法
go-mapper-gen generate

# 指定配置文件
./go-mapper-gen generate -c gobatis_config.yaml

# 指定输出目录
go-mapper-gen generate -o ./output

# 只生成结构体
go-mapper-gen generate --dao=false --sql=false

# 禁用 JSON 标签
go-mapper-gen generate --json-tag=false
```

### 使用 go:generate

在你的 Go 文件中添加 `//go:generate` 注释：

```go
//go:generate go-mapper-gen generate --driver sqlite --dsn test.db --output-dir generated --package model --include users,products
```

然后运行：

```bash
go generate
```

## 生成的代码结构

```
generated/
├── users.go           # 用户结构体
├── products.go        # 产品结构体
├── orders.go          # 订单结构体
├── dao/
│   ├── users_dao.go        # 用户 DAO 接口
│   ├── users_dao_impl.go   # 用户 DAO 实现
│   ├── products_dao.go     # 产品 DAO 接口
│   ├── products_dao_impl.go # 产品 DAO 实现
│   ├── orders_dao.go       # 订单 DAO 接口
│   └── orders_dao_impl.go  # 订单 DAO 实现
├── mapper/
│   ├── users_mapper.xml    # 用户 XML 映射文件
│   ├── products_mapper.xml # 产品 XML 映射文件
│   └── orders_mapper.xml   # 订单 XML 映射文件
└── sql/
    ├── users.sql       # 用户 SQL
    ├── products.sql    # 产品 SQL
    └── orders.sql      # 订单 SQL
```

## 生成的代码示例

### 结构体 (users.go)

```go
package model

import "time"

type Users struct {
	Id        *int64     `json:"id"`
	Username  *string    `json:"username"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
	Phone     *string    `json:"phone"`
	Avatar    *string    `json:"avatar"`
	Status    *int64     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Users) TableName() string {
	return "users"
}
```

### Gobatis DAO 接口示例

```go
package dao

import (
	"go-mapper-gen/generated"
	"github.com/gobatis/gobatis"
)

type UsersDAO interface {
	Create(user *generated.Users) error
	CreateBatch(users []*generated.Users) error
	GetByID(id int64) (*generated.Users, error)
	UpdateByID(user *generated.Users) error
	DeleteByID(id int64) error
	DeleteByIDs(ids []int64) error
	Exists(id int64) (bool, error)
	GetAll() ([]*generated.Users, error)
	GetByPage(offset, limit int) ([]*generated.Users, error)
	Count() (int64, error)
}
```

### SQL 文件 (sql/users.sql)

```sql
-- 查询所有记录
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at FROM users;

-- 根据ID查询
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at FROM users WHERE id = ?;

-- 插入记录
INSERT INTO users (username, email, password, phone, avatar, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- 批量插入
INSERT INTO users (username, email, password, phone, avatar, status, created_at, updated_at) VALUES 
(?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?);

-- 更新记录
UPDATE users SET username = ?, email = ?, password = ?, phone = ?, avatar = ?, status = ?, updated_at = ? WHERE id = ?;

-- 删除记录
DELETE FROM users WHERE id = ?;

-- 分页查询
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at FROM users LIMIT ? OFFSET ?;

-- 统计总数
SELECT COUNT(*) FROM users;
```

## 使用示例

### Gobatis 使用示例

```go
package main

import (
	"database/sql"
	"log"
	"go-mapper-gen/generated"
	"go-mapper-gen/generated/dao"
	"github.com/gobatis/gobatis"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 连接数据库
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	// 创建表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		phone TEXT,
		avatar TEXT,
		status INTEGER DEFAULT 1,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal("创建表失败:", err)
	}

	// 配置 Gobatis
	config := gobatis.NewConfig()
	config.SetDB(db)
	
	// 加载 XML 映射文件
	if err := config.LoadMapperFromFile("generated/mapper/users_mapper.xml"); err != nil {
		log.Fatal("加载映射文件失败:", err)
	}

	// 创建 Session
	session := config.NewSession()
	defer session.Close()

	// 创建 DAO 实例
	userDAO := &dao.UsersDAOImpl{Session: session}

	// 创建用户
	user := &generated.Users{
		Username: stringPtr("john_doe"),
		Email:    stringPtr("john@example.com"),
		Password: stringPtr("password123"),
	}

	if err := userDAO.Create(user); err != nil {
		log.Fatal("创建用户失败:", err)
	}

	// 查询所有用户
	users, err := userDAO.GetAll()
	if err != nil {
		log.Fatal("查询用户失败:", err)
	}

	log.Printf("找到 %d 个用户", len(users))
	for _, u := range users {
		log.Printf("用户: %+v", u)
	}

	// 统计用户数量
	count, err := userDAO.Count()
	if err != nil {
		log.Fatal("统计用户失败:", err)
	}
	log.Printf("用户总数: %d", count)
}

func stringPtr(s string) *string {
	return &s
}
```

完整的使用示例可以在 `examples/usage_example.go` 中找到：

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"go-mapper-gen/generated/dao"
	model "go-mapper-gen/generated"
)

func main() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 创建 DAO
	usersDAO := dao.NewUsersDAO(db)
	ctx := context.Background()

	// 查询所有用户
	users, err := usersDAO.GetAll(ctx)
	if err != nil {
		log.Fatal("查询用户失败:", err)
	}

	fmt.Printf("找到 %d 个用户:\n", len(users))
	for _, user := range users {
		fmt.Printf("- ID: %v, 用户名: %v, 邮箱: %v\n", user.Id, user.Username, user.Email)
	}

	// 创建新用户
	newUser := &model.Users{
		Username:  stringPtr("test_user"),
		Email:     stringPtr("test@example.com"),
		Password:  stringPtr("password123"),
		Status:    int64Ptr(1),
		CreatedAt: timePtr(time.Now()),
		UpdatedAt: timePtr(time.Now()),
	}

	err = usersDAO.Create(ctx, newUser)
	if err != nil {
		log.Fatal("创建用户失败:", err)
	}
	fmt.Printf("成功创建用户: %+v\n", newUser)

	// 分页查询
	pageUsers, err := usersDAO.GetByPage(ctx, 0, 2)
	if err != nil {
		log.Fatal("分页查询失败:", err)
	}

	// 统计总数
	count, err := usersDAO.Count(ctx)
	if err != nil {
		log.Fatal("统计失败:", err)
	}

	fmt.Printf("\n分页查询结果 (总数: %d):\n", count)
	for _, user := range pageUsers {
		fmt.Printf("- %v (%v)\n", *user.Username, *user.Email)
	}
}

// 辅助函数
func stringPtr(s string) *string { return &s }
func int64Ptr(i int64) *int64    { return &i }
func timePtr(t time.Time) *time.Time { return &t }
```

## 项目特性

- ✅ **多数据库支持**: 支持 SQLite、MySQL、PostgreSQL
- ✅ **灵活配置**: 支持 YAML 配置文件和命令行参数
- ✅ **完整代码生成**: 自动生成 Struct、DAO、SQL 文件
- ✅ **Gobatis 集成**: 生成的代码完全兼容 Gobatis 框架
- ✅ **类型安全**: 使用指针类型处理 NULL 值
- ✅ **模板化**: 基于 Go template 的代码生成
- ✅ **go:generate 支持**: 可以集成到 Go 工具链中
- ✅ **XML 映射**: 自动生成 Gobatis XML 映射文件
- ✅ **分页查询**: 内置分页查询支持
- ✅ **批量操作**: 支持批量插入和删除

## 项目结构

```
go-mapper-gen/
├── cmd/go-mapper-gen/          # 命令行入口
├── internal/
│   ├── cmd/                    # 命令行处理
│   ├── config/                 # 配置文件处理
│   ├── database/               # 数据库连接和元数据读取
│   └── generator/              # 代码生成器
├── examples/                   # 使用示例
├── generated/                  # 生成的代码输出目录
└── README.md
```

## 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 配置说明

详细配置选项请参考 [配置文档](docs/config.md)。

## 支持的数据库

- MySQL 5.7+
- PostgreSQL 10+
- SQLite 3+

## 开发

```bash
# 克隆项目
git clone https://github.com/your-username/go-mapper-gen.git
cd go-mapper-gen

# 安装依赖
go mod tidy

# 运行测试
go test ./...

# 构建
go build -o go-mapper-gen ./cmd/go-mapper-gen
```

## 许可证

MIT License