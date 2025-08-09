# Go Mapper Generator

A powerful Go code generation tool that automatically generates Go structs, DAO layer code, and SQL statements from database table structures. Supports MySQL, PostgreSQL, and SQLite databases, specifically designed for the Gobatis framework.

## Features

- 🚀 **Multi-Database Support**: MySQL, PostgreSQL, SQLite
- 📝 **Code Generation**: Automatically generate Go structs, DAO interfaces and implementations
- 🏷️ **Tag Support**: Automatic JSON tag generation
- 🔧 **Highly Configurable**: Support for table filtering, field mapping, naming conventions, etc.
- 📊 **SQL Generation**: Automatically generate common CRUD SQL statements
- 🎯 **Gobatis Support**: Specifically designed for the Gobatis framework, generating complete DAO interfaces, implementations, and XML mapping files

## Installation

Build from source:

```bash
git clone https://github.com/your-username/go-mapper-gen.git
cd go-mapper-gen
go build -o go-mapper-gen ./cmd/go-mapper-gen
```

Or run directly:

```bash
go run ./cmd/go-mapper-gen
```

## Usage

### 1. Configuration File Method

Create configuration file `generator.yaml`:

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
  generate_dao: true    # Generate DAO layer code
  generate_sql: true    # Generate SQL statements
  json_tag: true        # Generate JSON tags
  generate_example: true # Generate Example methods
  namespace_format: "{struct}DAO"  # XML namespace format, supports {struct} placeholder
```

Then run:

```bash
go-mapper-gen generate -c generator.yaml
```

### 2. Command Line Method

```bash
go-mapper-gen generate \
  --driver mysql \
  --dsn "user:password@tcp(localhost:3306)/database" \
  --output ./generated \
  --package model \
  --tables users,orders
```

### 3. go:generate Integration

Add to your Go file:

```go
//go:generate go-mapper-gen generate -c generator.yaml
```

Then run:

```bash
go generate
```

### Command Line Usage

```bash
# Generate code using configuration file
./go-mapper-gen generate -c generator-sqlite.yaml

# Generate code with direct parameters
./go-mapper-gen generate --driver sqlite --dsn test.db --output-dir generated --package model

# View help
./go-mapper-gen --help
```

### Command Line Parameters

```bash
# Basic usage
go-mapper-gen generate

# Specify configuration file
./go-mapper-gen generate -c gobatis_config.yaml

# Specify output directory
go-mapper-gen generate -o ./output

# Generate structs only
go-mapper-gen generate --dao=false --sql=false

# Disable JSON tags
go-mapper-gen generate --json-tag=false
```

### Using go:generate

Add `//go:generate` comment to your Go file:

```go
//go:generate go-mapper-gen generate --driver sqlite --dsn test.db --output-dir generated --package model --include users,products
```

Then run:

```bash
go generate
```

## Generated Code Structure

```
generated/
├── users.go           # User struct
├── products.go        # Product struct
├── orders.go          # Order struct
├── dao/
│   ├── users_dao.go        # User DAO interface
│   ├── users_dao_impl.go   # User DAO implementation
│   ├── products_dao.go     # Product DAO interface
│   ├── products_dao_impl.go # Product DAO implementation
│   ├── orders_dao.go       # Order DAO interface
│   └── orders_dao_impl.go  # Order DAO implementation
├── mapper/
│   ├── users_mapper.xml    # User XML mapping file
│   ├── products_mapper.xml # Product XML mapping file
│   └── orders_mapper.xml   # Order XML mapping file
└── sql/
    ├── users.sql       # User SQL
    ├── products.sql    # Product SQL
    └── orders.sql      # Order SQL
```

## Generated Code Examples

### Struct (users.go)

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

### Gobatis DAO Interface Example

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

### SQL File (sql/users.sql)

```sql
-- Query all records
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at FROM users;

-- Query by ID
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at FROM users WHERE id = ?;

-- Insert record
INSERT INTO users (username, email, password, phone, avatar, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- Batch insert
INSERT INTO users (username, email, password, phone, avatar, status, created_at, updated_at) VALUES 
(?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?);

-- Update record
UPDATE users SET username = ?, email = ?, password = ?, phone = ?, avatar = ?, status = ?, updated_at = ? WHERE id = ?;

-- Delete record
DELETE FROM users WHERE id = ?;

-- Paginated query
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at FROM users LIMIT ? OFFSET ?;

-- Count total
SELECT COUNT(*) FROM users;
```

## Usage Examples

### Gobatis Usage Example

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
	// Connect to database
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Create table
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
		log.Fatal("Failed to create table:", err)
	}

	// Configure Gobatis
	config := gobatis.NewConfig()
	config.SetDB(db)
	
	// Load XML mapping file
	if err := config.LoadMapperFromFile("generated/mapper/users_mapper.xml"); err != nil {
		log.Fatal("Failed to load mapping file:", err)
	}

	// Create Session
	session := config.NewSession()
	defer session.Close()

	// Create DAO instance
	userDAO := &dao.UsersDAOImpl{Session: session}

	// Create user
	user := &generated.Users{
		Username: stringPtr("john_doe"),
		Email:    stringPtr("john@example.com"),
		Password: stringPtr("password123"),
	}

	if err := userDAO.Create(user); err != nil {
		log.Fatal("Failed to create user:", err)
	}

	// Query all users
	users, err := userDAO.GetAll()
	if err != nil {
		log.Fatal("Failed to query users:", err)
	}

	log.Printf("Found %d users", len(users))
	for _, u := range users {
		log.Printf("User: %+v", u)
	}

	// Count users
	count, err := userDAO.Count()
	if err != nil {
		log.Fatal("Failed to count users:", err)
	}
	log.Printf("Total users: %d", count)
}

func stringPtr(s string) *string {
	return &s
}
```

Complete usage example can be found in `examples/usage_example.go`:

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
	// Connect to database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create DAO
	usersDAO := dao.NewUsersDAO(db)
	ctx := context.Background()

	// Query all users
	users, err := usersDAO.GetAll(ctx)
	if err != nil {
		log.Fatal("Failed to query users:", err)
	}

	fmt.Printf("Found %d users:\n", len(users))
	for _, user := range users {
		fmt.Printf("- ID: %v, Username: %v, Email: %v\n", user.Id, user.Username, user.Email)
	}

	// Create new user
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
		log.Fatal("Failed to create user:", err)
	}
	fmt.Printf("Successfully created user: %+v\n", newUser)

	// Paginated query
	pageUsers, err := usersDAO.GetByPage(ctx, 0, 2)
	if err != nil {
		log.Fatal("Failed to query with pagination:", err)
	}

	// Count total
	count, err := usersDAO.Count(ctx)
	if err != nil {
		log.Fatal("Failed to count:", err)
	}

	fmt.Printf("\nPaginated query results (total: %d):\n", count)
	for _, user := range pageUsers {
		fmt.Printf("- %v (%v)\n", *user.Username, *user.Email)
	}
}

// Helper functions
func stringPtr(s string) *string { return &s }
func int64Ptr(i int64) *int64    { return &i }
func timePtr(t time.Time) *time.Time { return &t }
```

## Project Features

- ✅ **Multi-Database Support**: Supports SQLite, MySQL, PostgreSQL
- ✅ **Flexible Configuration**: Supports YAML configuration files and command line parameters
- ✅ **Complete Code Generation**: Automatically generates Struct, DAO, SQL files
- ✅ **Gobatis Integration**: Generated code is fully compatible with Gobatis framework
- ✅ **Type Safety**: Uses pointer types to handle NULL values
- ✅ **Template-based**: Code generation based on Go templates
- ✅ **go:generate Support**: Can be integrated into Go toolchain
- ✅ **XML Mapping**: Automatically generates Gobatis XML mapping files
- ✅ **Paginated Queries**: Built-in pagination query support
- ✅ **Batch Operations**: Supports batch insert and delete

## Project Structure

```
go-mapper-gen/
├── cmd/go-mapper-gen/          # Command line entry point
├── internal/
│   ├── cmd/                    # Command line processing
│   ├── config/                 # Configuration file processing
│   ├── database/               # Database connection and metadata reading
│   └── generator/              # Code generators
├── examples/                   # Usage examples
├── generated/                  # Generated code output directory
└── README.md
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork this project
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Configuration Guide

### Configuration Options Details

#### Database Configuration
- `driver`: Database driver type (mysql, postgres, sqlite)
- `dsn`: Database connection string

#### Output Configuration
- `dir`: Code output directory
- `package`: Package name for generated code

#### Tables Configuration
- `include`: List of table names to include, empty means include all tables
- `exclude`: List of table names to exclude
- `prefix`: Table name prefix, will be removed when generating structs

#### Options Configuration
- `generate_dao`: Whether to generate DAO layer code (default: true)
- `generate_sql`: Whether to generate SQL files (default: true)
- `json_tag`: Whether to generate JSON tags (default: true)
- `generate_example`: Whether to generate Example methods (default: true)
- `namespace_format`: XML namespace format template (default: "{struct}DAO")

#### XML Namespace Customization

The `namespace_format` configuration option allows you to customize the namespace format in XML mapping files. Supports the following placeholders:

- `{struct}`: Will be replaced with the struct name

**Configuration Examples:**

```yaml
options:
  # Default format: UsersDAO, ProductsDAO
  namespace_format: "{struct}DAO"
  
  # Custom format: UsersMapper, ProductsMapper  
  namespace_format: "{struct}Mapper"
  
  # Format with package name: com.example.UsersDAO
  namespace_format: "com.example.{struct}DAO"
  
  # Struct name only: Users, Products
  namespace_format: "{struct}"
```

**Generated XML Examples:**

```xml
<!-- namespace_format: "{struct}DAO" -->
<mapper namespace="UsersDAO">

<!-- namespace_format: "{struct}Mapper" -->
<mapper namespace="UsersMapper">

<!-- namespace_format: "com.example.{struct}DAO" -->
<mapper namespace="com.example.UsersDAO">
```

For detailed configuration options, please refer to the [Configuration Documentation](docs/config.md).

## Supported Databases

- MySQL 5.7+
- PostgreSQL 10+
- SQLite 3+

## Development

```bash
# Clone the project
git clone https://github.com/your-username/go-mapper-gen.git
cd go-mapper-gen

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build
go build -o go-mapper-gen ./cmd/go-mapper-gen
```

## License

MIT License