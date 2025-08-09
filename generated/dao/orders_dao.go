package dao

import (
	"context"
	model "generated_gobatis/model"
)

// OrdersDAO Orders 数据访问接口
// 符合 Gobatis 框架规范的方法命名和参数定义
type OrdersDAO interface {
	// Insert 方法 - 插入操作
	// Insert 插入单个Orders记录
	Insert(ctx context.Context, record *model.Orders) error
	
	// InsertBatch 批量插入Orders记录
	InsertBatch(ctx context.Context, records []*model.Orders) error


	// Select 方法 - 查询操作
	// SelectById 根据主键查询Orders
	SelectById(ctx context.Context, id *int) (*model.Orders, error)
	
	// Update 方法 - 更新操作
	// UpdateById 根据主键更新Orders
	UpdateById(ctx context.Context, record *model.Orders) error
	
	// Delete 方法 - 删除操作
	// DeleteById 根据主键删除Orders
	DeleteById(ctx context.Context, id *int) error
	
	// DeleteByIds 根据主键列表批量删除Orders
	DeleteByIds(ctx context.Context, ids []*int) error


	// SelectAll 查询所有Orders记录
	SelectAll(ctx context.Context) ([]*model.Orders, error)
	
	// SelectByPage 分页查询Orders记录
	SelectByPage(ctx context.Context, offset, limit int) ([]*model.Orders, error)
	
	// Count 统计Orders记录总数
	Count(ctx context.Context) (int64, error)
	
	// SelectByCondition 根据条件动态查询Orders记录
	SelectByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.Orders, error)
	
	// CountByCondition 根据条件统计Orders记录数
	CountByCondition(ctx context.Context, condition map[string]interface{}) (int64, error)


	// ExistsById 检查指定主键的Orders记录是否存在
	ExistsById(ctx context.Context, id *int) (bool, error)


	// 兼容性方法 - 保持向后兼容
	// Create 创建Orders (兼容方法，内部调用 Insert)
	Create(ctx context.Context, orders *model.Orders) error

	// CreateBatch 批量创建Orders (兼容方法，内部调用 InsertBatch)
	CreateBatch(ctx context.Context, orderss []*model.Orders) error


	// GetByID 根据ID获取Orders (兼容方法，内部调用 SelectById)
	GetByID(ctx context.Context, id *int) (*model.Orders, error)

	// UpdateByID 根据ID更新Orders (兼容方法，内部调用 UpdateById)
	UpdateByID(ctx context.Context, orders *model.Orders) error

	// DeleteByID 根据ID删除Orders (兼容方法，内部调用 DeleteById)
	DeleteByID(ctx context.Context, id *int) error

	// DeleteByIDs 根据ID列表批量删除Orders (兼容方法，内部调用 DeleteByIds)
	DeleteByIDs(ctx context.Context, ids []*int) error

	// Exists 检查Orders是否存在 (兼容方法，内部调用 ExistsById)
	Exists(ctx context.Context, id *int) (bool, error)


	// GetAll 获取所有Orders (兼容方法，内部调用 SelectAll)
	GetAll(ctx context.Context) ([]*model.Orders, error)

	// GetByPage 分页获取Orders (兼容方法，内部调用 SelectByPage)
	GetByPage(ctx context.Context, offset, limit int) ([]*model.Orders, error)

	// FindByCondition 动态条件查询Orders (兼容方法，内部调用 SelectByCondition)
	FindByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.Orders, error)
}
