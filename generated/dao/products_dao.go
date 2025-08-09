package dao

import (
	"context"
	model "generated_gobatis/model"
)

// ProductsDAO Products 数据访问接口
// 符合 Gobatis 框架规范的方法命名和参数定义
type ProductsDAO interface {
	// Insert 方法 - 插入操作
	// Insert 插入单个Products记录
	Insert(ctx context.Context, record *model.Products) error
	
	// InsertBatch 批量插入Products记录
	InsertBatch(ctx context.Context, records []*model.Products) error


	// Select 方法 - 查询操作
	// SelectById 根据主键查询Products
	SelectById(ctx context.Context, id *int) (*model.Products, error)
	
	// Update 方法 - 更新操作
	// UpdateById 根据主键更新Products
	UpdateById(ctx context.Context, record *model.Products) error
	
	// Delete 方法 - 删除操作
	// DeleteById 根据主键删除Products
	DeleteById(ctx context.Context, id *int) error
	
	// DeleteByIds 根据主键列表批量删除Products
	DeleteByIds(ctx context.Context, ids []*int) error


	// SelectAll 查询所有Products记录
	SelectAll(ctx context.Context) ([]*model.Products, error)
	
	// SelectByPage 分页查询Products记录
	SelectByPage(ctx context.Context, offset, limit int) ([]*model.Products, error)
	
	// Count 统计Products记录总数
	Count(ctx context.Context) (int64, error)
	
	// SelectByCondition 根据条件动态查询Products记录
	SelectByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.Products, error)
	
	// CountByCondition 根据条件统计Products记录数
	CountByCondition(ctx context.Context, condition map[string]interface{}) (int64, error)


	// ExistsById 检查指定主键的Products记录是否存在
	ExistsById(ctx context.Context, id *int) (bool, error)


	// 兼容性方法 - 保持向后兼容
	// Create 创建Products (兼容方法，内部调用 Insert)
	Create(ctx context.Context, products *model.Products) error

	// CreateBatch 批量创建Products (兼容方法，内部调用 InsertBatch)
	CreateBatch(ctx context.Context, productss []*model.Products) error


	// GetByID 根据ID获取Products (兼容方法，内部调用 SelectById)
	GetByID(ctx context.Context, id *int) (*model.Products, error)

	// UpdateByID 根据ID更新Products (兼容方法，内部调用 UpdateById)
	UpdateByID(ctx context.Context, products *model.Products) error

	// DeleteByID 根据ID删除Products (兼容方法，内部调用 DeleteById)
	DeleteByID(ctx context.Context, id *int) error

	// DeleteByIDs 根据ID列表批量删除Products (兼容方法，内部调用 DeleteByIds)
	DeleteByIDs(ctx context.Context, ids []*int) error

	// Exists 检查Products是否存在 (兼容方法，内部调用 ExistsById)
	Exists(ctx context.Context, id *int) (bool, error)


	// GetAll 获取所有Products (兼容方法，内部调用 SelectAll)
	GetAll(ctx context.Context) ([]*model.Products, error)

	// GetByPage 分页获取Products (兼容方法，内部调用 SelectByPage)
	GetByPage(ctx context.Context, offset, limit int) ([]*model.Products, error)

	// FindByCondition 动态条件查询Products (兼容方法，内部调用 SelectByCondition)
	FindByCondition(ctx context.Context, condition map[string]interface{}) ([]*model.Products, error)
}
