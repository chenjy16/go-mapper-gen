package dao

import (
	model "go-mapper-gen/examples/generated/model"
	"gobatis/core/example"
)

// ProductsDAO Products 数据访问接口
// 严格遵循 GoBatis 框架方法命名规则和返回值规范
type ProductsDAO interface {
	// 插入方法 (INSERT) - 返回影响行数
	// Insert 插入单个Products记录
	Insert(record *model.Products) (int64, error)
	
	// InsertBatch 批量插入Products记录
	InsertBatch(records []*model.Products) (int64, error)
	
	// Add 添加Products记录 (Insert 的别名)
	Add(record *model.Products) (int64, error)
	
	// Create 创建Products记录 (Insert 的别名)
	Create(record *model.Products) (int64, error)
	
	// Save 保存Products记录 (Insert 的别名)
	Save(record *model.Products) (int64, error)


	// 查询方法 (SELECT) - 返回查询结果
	// GetById 根据主键获取Products
	GetById(id *int) (*model.Products, error)
	
	// FindById 根据主键查找Products (GetById 的别名)
	FindById(id *int) (*model.Products, error)
	
	// SelectById 根据主键选择Products (GetById 的别名)
	SelectById(id *int) (*model.Products, error)


	// GetAll 获取所有Products记录
	GetAll() ([]*model.Products, error)
	
	// FindAll 查找所有Products记录 (GetAll 的别名)
	FindAll() ([]*model.Products, error)
	
	// SelectAll 选择所有Products记录 (GetAll 的别名)
	SelectAll() ([]*model.Products, error)
	
	// ListAll 列出所有Products记录 (GetAll 的别名)
	ListAll() ([]*model.Products, error)
	
	// QueryAll 查询所有Products记录 (GetAll 的别名)
	QueryAll() ([]*model.Products, error)
	
	// GetByPage 分页获取Products记录
	GetByPage(offset, limit int) ([]*model.Products, error)
	
	// FindByPage 分页查找Products记录 (GetByPage 的别名)
	FindByPage(offset, limit int) ([]*model.Products, error)
	
	// SelectByPage 分页选择Products记录 (GetByPage 的别名)
	SelectByPage(offset, limit int) ([]*model.Products, error)
	
	// GetByCondition 根据条件获取Products记录
	GetByCondition(condition map[string]interface{}) ([]*model.Products, error)
	
	// FindByCondition 根据条件查找Products记录 (GetByCondition 的别名)
	FindByCondition(condition map[string]interface{}) ([]*model.Products, error)
	
	// SelectByCondition 根据条件选择Products记录 (GetByCondition 的别名)
	SelectByCondition(condition map[string]interface{}) ([]*model.Products, error)
	
	// QueryByCondition 根据条件查询Products记录 (GetByCondition 的别名)
	QueryByCondition(condition map[string]interface{}) ([]*model.Products, error)

	// 统计方法 - 返回数量
	// GetCount 获取Products记录总数
	GetCount() (int64, error)
	
	// Count 统计Products记录总数 (GetCount 的别名)
	Count() (int64, error)
	
	// CountByCondition 根据条件统计Products记录数
	CountByCondition(condition map[string]interface{}) (int64, error)


	// 存在性检查方法 - 返回布尔值
	// GetExistsById 检查指定主键的Products记录是否存在
	GetExistsById(id *int) (bool, error)


	// 更新方法 (UPDATE) - 返回影响行数

	// UpdateById 根据主键更新Products
	UpdateById(record *model.Products) (int64, error)
	
	// ModifyById 根据主键修改Products (UpdateById 的别名)
	ModifyById(record *model.Products) (int64, error)
	
	// EditById 根据主键编辑Products (UpdateById 的别名)
	EditById(record *model.Products) (int64, error)


	// UpdateByCondition 根据条件更新Products记录
	UpdateByCondition(record *model.Products, condition map[string]interface{}) (int64, error)

	// 删除方法 (DELETE) - 返回影响行数

	// DeleteById 根据主键删除Products
	DeleteById(id *int) (int64, error)
	
	// RemoveById 根据主键移除Products (DeleteById 的别名)
	RemoveById(id *int) (int64, error)
	
	// DeleteByIds 根据主键列表批量删除Products
	DeleteByIds(ids []*int) (int64, error)
	
	// RemoveByIds 根据主键列表批量移除Products (DeleteByIds 的别名)
	RemoveByIds(ids []*int) (int64, error)


	// DeleteByCondition 根据条件删除Products记录
	DeleteByCondition(condition map[string]interface{}) (int64, error)
	
	// RemoveByCondition 根据条件移除Products记录 (DeleteByCondition 的别名)
	RemoveByCondition(condition map[string]interface{}) (int64, error)


	// Example 查询方法 - 支持 GoBatis Example 功能
	// GetByExample 根据 Example 条件获取Products记录
	GetByExample(example *example.Example) ([]*model.Products, error)
	
	// FindByExample 根据 Example 条件查找Products记录 (GetByExample 的别名)
	FindByExample(example *example.Example) ([]*model.Products, error)
	
	// SelectByExample 根据 Example 条件选择Products记录 (GetByExample 的别名)
	SelectByExample(example *example.Example) ([]*model.Products, error)
	
	// QueryByExample 根据 Example 条件查询Products记录 (GetByExample 的别名)
	QueryByExample(example *example.Example) ([]*model.Products, error)
	
	// ListByExample 根据 Example 条件列出Products记录 (GetByExample 的别名)
	ListByExample(example *example.Example) ([]*model.Products, error)
	
	// CountByExample 根据 Example 条件统计Products记录数
	CountByExample(example *example.Example) (int64, error)
	
	// UpdateByExample 根据 Example 条件更新Products记录
	UpdateByExample(record *model.Products, example *example.Example) (int64, error)
	
	// ModifyByExample 根据 Example 条件修改Products记录 (UpdateByExample 的别名)
	ModifyByExample(record *model.Products, example *example.Example) (int64, error)
	
	// EditByExample 根据 Example 条件编辑Products记录 (UpdateByExample 的别名)
	EditByExample(record *model.Products, example *example.Example) (int64, error)
	
	// DeleteByExample 根据 Example 条件删除Products记录
	DeleteByExample(example *example.Example) (int64, error)
	
	// RemoveByExample 根据 Example 条件移除Products记录 (DeleteByExample 的别名)
	RemoveByExample(example *example.Example) (int64, error)

}
