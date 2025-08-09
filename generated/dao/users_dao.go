package dao

import (
	"context"
	model "generated_gobatis/model"
)

// UsersDAO Users 数据访问接口
// 符合 Gobatis 框架规范的方法命名和参数定义
type UsersDAO interface {
	// Insert 方法 - 插入操作
	// Insert 插入单个Users记录
	Insert(ctx context.Context, record *model.Users) error
	
	// InsertBatch 批量插入Users记录
	InsertBatch(ctx context.Context, records []*model.Users) error


	// Select 方法 - 查询操作
	// SelectById 根据主键查询Users
	SelectById(ctx context.Context, id *int) (*model.Users, error)
	
	// Update 方法 - 更新操作
	// UpdateById 根据主键更新Users
	UpdateById(ctx context.Context, record *model.Users) error
	
	// Delete 方法 - 删除操作
	// DeleteById 根据主键删除Users
	DeleteById(ctx context.Context, id *int) error
	
	// DeleteByIds 根据主键列表批量删除Users
	DeleteByIds(ctx context.Context, ids []*int) error


	// SelectAll 查询所有Users记录
	SelectAll(ctx context.Context) ([]*model.Users, error)
	
	// SelectByPage 分页查询Users记录
	SelectByPage(ctx context.Context, offset, limit int) ([]*model.Users, error)
	
	// Count 统计Users记录总数
	Count(ctx context.Context) (int64, error)
	
	// SelectByCondition 根据条件动态查询Users记录
	SelectByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.Users, error)
	
	// CountByCondition 根据条件统计Users记录数
	CountByCondition(ctx context.Context, condition map[string]interface{}) (int64, error)


	// ExistsById 检查指定主键的Users记录是否存在
	ExistsById(ctx context.Context, id *int) (bool, error)


	// 兼容性方法 - 保持向后兼容
	// Create 创建Users (兼容方法，内部调用 Insert)
	Create(ctx context.Context, users *model.Users) error

	// CreateBatch 批量创建Users (兼容方法，内部调用 InsertBatch)
	CreateBatch(ctx context.Context, userss []*model.Users) error


	// GetByID 根据ID获取Users (兼容方法，内部调用 SelectById)
	GetByID(ctx context.Context, id *int) (*model.Users, error)

	// UpdateByID 根据ID更新Users (兼容方法，内部调用 UpdateById)
	UpdateByID(ctx context.Context, users *model.Users) error

	// DeleteByID 根据ID删除Users (兼容方法，内部调用 DeleteById)
	DeleteByID(ctx context.Context, id *int) error

	// DeleteByIDs 根据ID列表批量删除Users (兼容方法，内部调用 DeleteByIds)
	DeleteByIDs(ctx context.Context, ids []*int) error

	// Exists 检查Users是否存在 (兼容方法，内部调用 ExistsById)
	Exists(ctx context.Context, id *int) (bool, error)


	// GetAll 获取所有Users (兼容方法，内部调用 SelectAll)
	GetAll(ctx context.Context) ([]*model.Users, error)

	// GetByPage 分页获取Users (兼容方法，内部调用 SelectByPage)
	GetByPage(ctx context.Context, offset, limit int) ([]*model.Users, error)

	// FindByCondition 动态条件查询Users (兼容方法，内部调用 SelectByCondition)
	FindByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.Users, error)
}
