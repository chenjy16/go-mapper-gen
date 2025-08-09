package dao

import (
	model "../model"
	"github.com/chenjy16/gobatis/core/example"
)

// OrdersDAO Orders 数据访问接口
// 严格遵循 GoBatis 框架方法命名规则和返回值规范
type OrdersDAO interface {
	// 插入方法 (INSERT) - 返回影响行数
	// Insert 插入单个Orders记录
	Insert(record *model.Orders) (int64, error)
	
	// InsertBatch 批量插入Orders记录
	InsertBatch(records []*model.Orders) (int64, error)
	
	// Add 添加Orders记录 (Insert 的别名)
	Add(record *model.Orders) (int64, error)
	
	// Create 创建Orders记录 (Insert 的别名)
	Create(record *model.Orders) (int64, error)
	
	// Save 保存Orders记录 (Insert 的别名)
	Save(record *model.Orders) (int64, error)


	// 查询方法 (SELECT) - 返回查询结果
	// GetById 根据主键获取Orders
	GetById(id *int) (*model.Orders, error)
	
	// FindById 根据主键查找Orders (GetById 的别名)
	FindById(id *int) (*model.Orders, error)
	
	// SelectById 根据主键选择Orders (GetById 的别名)
	SelectById(id *int) (*model.Orders, error)


	// GetAll 获取所有Orders记录
	GetAll() ([]*model.Orders, error)
	
	// FindAll 查找所有Orders记录 (GetAll 的别名)
	FindAll() ([]*model.Orders, error)
	
	// SelectAll 选择所有Orders记录 (GetAll 的别名)
	SelectAll() ([]*model.Orders, error)
	
	// ListAll 列出所有Orders记录 (GetAll 的别名)
	ListAll() ([]*model.Orders, error)
	
	// QueryAll 查询所有Orders记录 (GetAll 的别名)
	QueryAll() ([]*model.Orders, error)
	
	// GetByPage 分页获取Orders记录
	GetByPage(offset, limit int) ([]*model.Orders, error)
	
	// FindByPage 分页查找Orders记录 (GetByPage 的别名)
	FindByPage(offset, limit int) ([]*model.Orders, error)
	
	// SelectByPage 分页选择Orders记录 (GetByPage 的别名)
	SelectByPage(offset, limit int) ([]*model.Orders, error)
	
	// GetByCondition 根据条件获取Orders记录
	GetByCondition(condition map[string]interface{}) ([]*model.Orders, error)
	
	// FindByCondition 根据条件查找Orders记录 (GetByCondition 的别名)
	FindByCondition(condition map[string]interface{}) ([]*model.Orders, error)
	
	// SelectByCondition 根据条件选择Orders记录 (GetByCondition 的别名)
	SelectByCondition(condition map[string]interface{}) ([]*model.Orders, error)
	
	// QueryByCondition 根据条件查询Orders记录 (GetByCondition 的别名)
	QueryByCondition(condition map[string]interface{}) ([]*model.Orders, error)

	// 统计方法 - 返回数量
	// GetCount 获取Orders记录总数
	GetCount() (int64, error)
	
	// Count 统计Orders记录总数 (GetCount 的别名)
	Count() (int64, error)
	
	// CountByCondition 根据条件统计Orders记录数
	CountByCondition(condition map[string]interface{}) (int64, error)


	// 存在性检查方法 - 返回布尔值
	// GetExistsById 检查指定主键的Orders记录是否存在
	GetExistsById(id *int) (bool, error)


	// 更新方法 (UPDATE) - 返回影响行数

	// UpdateById 根据主键更新Orders
	UpdateById(record *model.Orders) (int64, error)
	
	// ModifyById 根据主键修改Orders (UpdateById 的别名)
	ModifyById(record *model.Orders) (int64, error)
	
	// EditById 根据主键编辑Orders (UpdateById 的别名)
	EditById(record *model.Orders) (int64, error)


	// UpdateByCondition 根据条件更新Orders记录
	UpdateByCondition(record *model.Orders, condition map[string]interface{}) (int64, error)

	// 删除方法 (DELETE) - 返回影响行数

	// DeleteById 根据主键删除Orders
	DeleteById(id *int) (int64, error)
	
	// RemoveById 根据主键移除Orders (DeleteById 的别名)
	RemoveById(id *int) (int64, error)
	
	// DeleteByIds 根据主键列表批量删除Orders
	DeleteByIds(ids []*int) (int64, error)
	
	// RemoveByIds 根据主键列表批量移除Orders (DeleteByIds 的别名)
	RemoveByIds(ids []*int) (int64, error)


	// DeleteByCondition 根据条件删除Orders记录
	DeleteByCondition(condition map[string]interface{}) (int64, error)
	
	// RemoveByCondition 根据条件移除Orders记录 (DeleteByCondition 的别名)
	RemoveByCondition(condition map[string]interface{}) (int64, error)


	// Example 查询方法 - 支持 GoBatis Example 功能
	// GetByExample 根据 Example 条件获取Orders记录
	GetByExample(example *example.Example) ([]*model.Orders, error)
	
	// FindByExample 根据 Example 条件查找Orders记录 (GetByExample 的别名)
	FindByExample(example *example.Example) ([]*model.Orders, error)
	
	// SelectByExample 根据 Example 条件选择Orders记录 (GetByExample 的别名)
	SelectByExample(example *example.Example) ([]*model.Orders, error)
	
	// QueryByExample 根据 Example 条件查询Orders记录 (GetByExample 的别名)
	QueryByExample(example *example.Example) ([]*model.Orders, error)
	
	// ListByExample 根据 Example 条件列出Orders记录 (GetByExample 的别名)
	ListByExample(example *example.Example) ([]*model.Orders, error)
	
	// CountByExample 根据 Example 条件统计Orders记录数
	CountByExample(example *example.Example) (int64, error)
	
	// UpdateByExample 根据 Example 条件更新Orders记录
	UpdateByExample(record *model.Orders, example *example.Example) (int64, error)
	
	// ModifyByExample 根据 Example 条件修改Orders记录 (UpdateByExample 的别名)
	ModifyByExample(record *model.Orders, example *example.Example) (int64, error)
	
	// EditByExample 根据 Example 条件编辑Orders记录 (UpdateByExample 的别名)
	EditByExample(record *model.Orders, example *example.Example) (int64, error)
	
	// DeleteByExample 根据 Example 条件删除Orders记录
	DeleteByExample(example *example.Example) (int64, error)
	
	// RemoveByExample 根据 Example 条件移除Orders记录 (DeleteByExample 的别名)
	RemoveByExample(example *example.Example) (int64, error)

}
