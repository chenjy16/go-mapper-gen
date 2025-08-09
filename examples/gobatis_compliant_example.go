package main

import (
	"context"
	"fmt"
	"log"
	"time"

	// 注意：实际使用时需要导入生成的代码包
	// model "./generated/model"
	// dao "./generated/dao"
	// "github.com/runner-mei/GoBatis"
	// "github.com/runner-mei/GoBatis/example"
)

// 模拟生成的 Users 结构体和 DAO 接口
// 注意：实际使用时应该导入生成的代码

// Users 模拟生成的用户模型结构体
type Users struct {
	Id       *int    `db:"id" json:"id"`
	Username *string `db:"username" json:"username"`
	Email    *string `db:"email" json:"email"`
	Password *string `db:"password" json:"password"`
}

// 模拟 GoBatis 会话
type MockSession struct{}

func (s *MockSession) SelectOne(statementID string, params interface{}) (interface{}, error) {
	fmt.Printf("执行 SelectOne: %s, 参数: %+v\n", statementID, params)
	return &Users{Id: intPtr(1), Username: stringPtr("john_doe"), Email: stringPtr("john@example.com")}, nil
}

func (s *MockSession) SelectList(statementID string, params interface{}) ([]interface{}, error) {
	fmt.Printf("执行 SelectList: %s, 参数: %+v\n", statementID, params)
	return []interface{}{
		&Users{Id: intPtr(1), Username: stringPtr("john_doe"), Email: stringPtr("john@example.com")},
		&Users{Id: intPtr(2), Username: stringPtr("jane_doe"), Email: stringPtr("jane@example.com")},
	}, nil
}

func (s *MockSession) Insert(statementID string, params interface{}) (int64, error) {
	fmt.Printf("执行 Insert: %s, 参数: %+v\n", statementID, params)
	return 1, nil
}

func (s *MockSession) Update(statementID string, params interface{}) (int64, error) {
	fmt.Printf("执行 Update: %s, 参数: %+v\n", statementID, params)
	return 1, nil
}

func (s *MockSession) Delete(statementID string, params interface{}) (int64, error) {
	fmt.Printf("执行 Delete: %s, 参数: %+v\n", statementID, params)
	return 1, nil
}

// 模拟 DAO 实现（实际由 GoBatis 代理生成）
type MockUsersDAO struct {
	session *MockSession
}

// 实现符合 GoBatis 规则的方法

// 插入方法 (INSERT) - 返回影响行数
func (dao *MockUsersDAO) Insert(ctx context.Context, record *Users) (int64, error) {
	return dao.session.Insert("UsersDAO.Insert", record)
}

func (dao *MockUsersDAO) Add(ctx context.Context, record *Users) (int64, error) {
	return dao.session.Insert("UsersDAO.Add", record)
}

func (dao *MockUsersDAO) Create(ctx context.Context, record *Users) (int64, error) {
	return dao.session.Insert("UsersDAO.Create", record)
}

func (dao *MockUsersDAO) Save(ctx context.Context, record *Users) (int64, error) {
	return dao.session.Insert("UsersDAO.Save", record)
}

// 查询方法 (SELECT) - 返回查询结果
func (dao *MockUsersDAO) GetById(ctx context.Context, id *int) (*Users, error) {
	result, err := dao.session.SelectOne("UsersDAO.GetById", id)
	if err != nil {
		return nil, err
	}
	return result.(*Users), nil
}

func (dao *MockUsersDAO) FindById(ctx context.Context, id *int) (*Users, error) {
	result, err := dao.session.SelectOne("UsersDAO.FindById", id)
	if err != nil {
		return nil, err
	}
	return result.(*Users), nil
}

func (dao *MockUsersDAO) SelectById(ctx context.Context, id *int) (*Users, error) {
	result, err := dao.session.SelectOne("UsersDAO.SelectById", id)
	if err != nil {
		return nil, err
	}
	return result.(*Users), nil
}

func (dao *MockUsersDAO) GetAll(ctx context.Context) ([]*Users, error) {
	results, err := dao.session.SelectList("UsersDAO.GetAll", nil)
	if err != nil {
		return nil, err
	}
	users := make([]*Users, len(results))
	for i, result := range results {
		users[i] = result.(*Users)
	}
	return users, nil
}

func (dao *MockUsersDAO) FindAll(ctx context.Context) ([]*Users, error) {
	results, err := dao.session.SelectList("UsersDAO.FindAll", nil)
	if err != nil {
		return nil, err
	}
	users := make([]*Users, len(results))
	for i, result := range results {
		users[i] = result.(*Users)
	}
	return users, nil
}

func (dao *MockUsersDAO) SelectAll(ctx context.Context) ([]*Users, error) {
	results, err := dao.session.SelectList("UsersDAO.SelectAll", nil)
	if err != nil {
		return nil, err
	}
	users := make([]*Users, len(results))
	for i, result := range results {
		users[i] = result.(*Users)
	}
	return users, nil
}

func (dao *MockUsersDAO) QueryAll(ctx context.Context) ([]*Users, error) {
	results, err := dao.session.SelectList("UsersDAO.QueryAll", nil)
	if err != nil {
		return nil, err
	}
	users := make([]*Users, len(results))
	for i, result := range results {
		users[i] = result.(*Users)
	}
	return users, nil
}

func (dao *MockUsersDAO) ListAll(ctx context.Context) ([]*Users, error) {
	results, err := dao.session.SelectList("UsersDAO.ListAll", nil)
	if err != nil {
		return nil, err
	}
	users := make([]*Users, len(results))
	for i, result := range results {
		users[i] = result.(*Users)
	}
	return users, nil
}

// 更新方法 (UPDATE) - 返回影响行数
func (dao *MockUsersDAO) UpdateById(ctx context.Context, record *Users) (int64, error) {
	return dao.session.Update("UsersDAO.UpdateById", record)
}

func (dao *MockUsersDAO) ModifyById(ctx context.Context, record *Users) (int64, error) {
	return dao.session.Update("UsersDAO.ModifyById", record)
}

func (dao *MockUsersDAO) EditById(ctx context.Context, record *Users) (int64, error) {
	return dao.session.Update("UsersDAO.EditById", record)
}

// 删除方法 (DELETE) - 返回影响行数
func (dao *MockUsersDAO) DeleteById(ctx context.Context, id *int) (int64, error) {
	return dao.session.Delete("UsersDAO.DeleteById", id)
}

func (dao *MockUsersDAO) RemoveById(ctx context.Context, id *int) (int64, error) {
	return dao.session.Delete("UsersDAO.RemoveById", id)
}

// 统计方法
func (dao *MockUsersDAO) Count(ctx context.Context) (int64, error) {
	// 模拟统计
	fmt.Println("执行 Count: UsersDAO.Count")
	return 10, nil
}

func (dao *MockUsersDAO) GetCount(ctx context.Context) (int64, error) {
	// 模拟统计
	fmt.Println("执行 GetCount: UsersDAO.GetCount")
	return 10, nil
}

// 辅助函数
func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func main() {
	fmt.Println("=== GoBatis 方法命名规则符合性示例 ===")
	
	ctx := context.Background()
	session := &MockSession{}
	userDAO := &MockUsersDAO{session: session}
	
	// 1. 插入方法示例 (INSERT) - 返回影响行数
	fmt.Println("\n1. 插入方法 (INSERT) - 返回影响行数:")
	newUser := &Users{
		Username: stringPtr("new_user"),
		Email:    stringPtr("new@example.com"),
		Password: stringPtr("password123"),
	}
	
	// 使用不同的插入方法前缀
	if rowsAffected, err := userDAO.Insert(ctx, newUser); err != nil {
		log.Printf("Insert 失败: %v", err)
	} else {
		fmt.Printf("✅ Insert 成功，影响行数: %d\n", rowsAffected)
	}
	
	if rowsAffected, err := userDAO.Add(ctx, newUser); err != nil {
		log.Printf("Add 失败: %v", err)
	} else {
		fmt.Printf("✅ Add 成功，影响行数: %d\n", rowsAffected)
	}
	
	if rowsAffected, err := userDAO.Create(ctx, newUser); err != nil {
		log.Printf("Create 失败: %v", err)
	} else {
		fmt.Printf("✅ Create 成功，影响行数: %d\n", rowsAffected)
	}
	
	if rowsAffected, err := userDAO.Save(ctx, newUser); err != nil {
		log.Printf("Save 失败: %v", err)
	} else {
		fmt.Printf("✅ Save 成功，影响行数: %d\n", rowsAffected)
	}
	
	// 2. 查询方法示例 (SELECT) - 返回查询结果
	fmt.Println("\n2. 查询方法 (SELECT) - 返回查询结果:")
	
	// 单个查询 - 使用不同的查询方法前缀
	id := intPtr(1)
	
	if user, err := userDAO.GetById(ctx, id); err != nil {
		log.Printf("GetById 失败: %v", err)
	} else {
		fmt.Printf("✅ GetById 成功: %+v\n", user)
	}
	
	if user, err := userDAO.FindById(ctx, id); err != nil {
		log.Printf("FindById 失败: %v", err)
	} else {
		fmt.Printf("✅ FindById 成功: %+v\n", user)
	}
	
	if user, err := userDAO.SelectById(ctx, id); err != nil {
		log.Printf("SelectById 失败: %v", err)
	} else {
		fmt.Printf("✅ SelectById 成功: %+v\n", user)
	}
	
	// 列表查询 - 使用不同的查询方法前缀
	if users, err := userDAO.GetAll(ctx); err != nil {
		log.Printf("GetAll 失败: %v", err)
	} else {
		fmt.Printf("✅ GetAll 成功，获取到 %d 个用户\n", len(users))
	}
	
	if users, err := userDAO.FindAll(ctx); err != nil {
		log.Printf("FindAll 失败: %v", err)
	} else {
		fmt.Printf("✅ FindAll 成功，获取到 %d 个用户\n", len(users))
	}
	
	if users, err := userDAO.SelectAll(ctx); err != nil {
		log.Printf("SelectAll 失败: %v", err)
	} else {
		fmt.Printf("✅ SelectAll 成功，获取到 %d 个用户\n", len(users))
	}
	
	if users, err := userDAO.QueryAll(ctx); err != nil {
		log.Printf("QueryAll 失败: %v", err)
	} else {
		fmt.Printf("✅ QueryAll 成功，获取到 %d 个用户\n", len(users))
	}
	
	if users, err := userDAO.ListAll(ctx); err != nil {
		log.Printf("ListAll 失败: %v", err)
	} else {
		fmt.Printf("✅ ListAll 成功，获取到 %d 个用户\n", len(users))
	}
	
	// 3. 更新方法示例 (UPDATE) - 返回影响行数
	fmt.Println("\n3. 更新方法 (UPDATE) - 返回影响行数:")
	updateUser := &Users{
		Id:       intPtr(1),
		Username: stringPtr("updated_user"),
		Email:    stringPtr("updated@example.com"),
	}
	
	if rowsAffected, err := userDAO.UpdateById(ctx, updateUser); err != nil {
		log.Printf("UpdateById 失败: %v", err)
	} else {
		fmt.Printf("✅ UpdateById 成功，影响行数: %d\n", rowsAffected)
	}
	
	if rowsAffected, err := userDAO.ModifyById(ctx, updateUser); err != nil {
		log.Printf("ModifyById 失败: %v", err)
	} else {
		fmt.Printf("✅ ModifyById 成功，影响行数: %d\n", rowsAffected)
	}
	
	if rowsAffected, err := userDAO.EditById(ctx, updateUser); err != nil {
		log.Printf("EditById 失败: %v", err)
	} else {
		fmt.Printf("✅ EditById 成功，影响行数: %d\n", rowsAffected)
	}
	
	// 4. 删除方法示例 (DELETE) - 返回影响行数
	fmt.Println("\n4. 删除方法 (DELETE) - 返回影响行数:")
	
	if rowsAffected, err := userDAO.DeleteById(ctx, id); err != nil {
		log.Printf("DeleteById 失败: %v", err)
	} else {
		fmt.Printf("✅ DeleteById 成功，影响行数: %d\n", rowsAffected)
	}
	
	if rowsAffected, err := userDAO.RemoveById(ctx, id); err != nil {
		log.Printf("RemoveById 失败: %v", err)
	} else {
		fmt.Printf("✅ RemoveById 成功，影响行数: %d\n", rowsAffected)
	}
	
	// 5. 统计方法示例
	fmt.Println("\n5. 统计方法:")
	
	if count, err := userDAO.Count(ctx); err != nil {
		log.Printf("Count 失败: %v", err)
	} else {
		fmt.Printf("✅ Count 成功，总数: %d\n", count)
	}
	
	if count, err := userDAO.GetCount(ctx); err != nil {
		log.Printf("GetCount 失败: %v", err)
	} else {
		fmt.Printf("✅ GetCount 成功，总数: %d\n", count)
	}
	
	fmt.Println("\n=== GoBatis 方法命名规则说明 ===")
	fmt.Println("✅ 查询方法 (SELECT): Get*, Find*, Select*, Query*, List*")
	fmt.Println("✅ 插入方法 (INSERT): Insert*, Add*, Create*, Save* - 返回 (int64, error)")
	fmt.Println("✅ 更新方法 (UPDATE): Update*, Modify*, Edit* - 返回 (int64, error)")
	fmt.Println("✅ 删除方法 (DELETE): Delete*, Remove* - 返回 (int64, error)")
	fmt.Println("✅ 返回值规则:")
	fmt.Println("   - 查询方法: (result, error) 或 ([]result, error)")
	fmt.Println("   - 增删改方法: (int64, error) 表示影响行数")
	fmt.Println("✅ Statement ID 格式: {InterfaceName}.{MethodName}")
	fmt.Println("\n=== 示例完成 ===")
}