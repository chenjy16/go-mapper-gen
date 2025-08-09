package dao

import (
	model "../model"
	"github.com/chenjy16/gobatis/core/example"
)

// OrderItemsDAO OrderItems 数据访问接口
// 严格遵循 GoBatis 框架方法命名规则和返回值规范
type OrderItemsDAO interface {
	// 插入方法 (INSERT) - 返回影响行数
	// Insert 插入单个OrderItems记录
	Insert(record *model.OrderItems) (int64, error)
	
	// InsertBatch 批量插入OrderItems记录
	InsertBatch(records []*model.OrderItems) (int64, error)
	
	// Add 添加OrderItems记录 (Insert 的别名)
	Add(record *model.OrderItems) (int64, error)
	
	// Create 创建OrderItems记录 (Insert 的别名)
	Create(record *model.OrderItems) (int64, error)
	
	// Save 保存OrderItems记录 (Insert 的别名)
	Save(record *model.OrderItems) (int64, error)


	// 查询方法 (SELECT) - 返回查询结果
	// GetById 根据主键获取OrderItems
	GetById(id *int) (*model.OrderItems, error)
	
	// FindById 根据主键查找OrderItems (GetById 的别名)
	FindById(id *int) (*model.OrderItems, error)
	
	// SelectById 根据主键选择OrderItems (GetById 的别名)
	SelectById(id *int) (*model.OrderItems, error)


	// GetAll 获取所有OrderItems记录
	GetAll() ([]*model.OrderItems, error)
	
	// FindAll 查找所有OrderItems记录 (GetAll 的别名)
	FindAll() ([]*model.OrderItems, error)
	
	// SelectAll 选择所有OrderItems记录 (GetAll 的别名)
	SelectAll() ([]*model.OrderItems, error)
	
	// ListAll 列出所有OrderItems记录 (GetAll 的别名)
	ListAll() ([]*model.OrderItems, error)
	
	// QueryAll 查询所有OrderItems记录 (GetAll 的别名)
	QueryAll() ([]*model.OrderItems, error)
	
	// GetByPage 分页获取OrderItems记录
	GetByPage(offset, limit int) ([]*model.OrderItems, error)
	
	// FindByPage 分页查找OrderItems记录 (GetByPage 的别名)
	FindByPage(offset, limit int) ([]*model.OrderItems, error)
	
	// SelectByPage 分页选择OrderItems记录 (GetByPage 的别名)
	SelectByPage(offset, limit int) ([]*model.OrderItems, error)
	
	// GetByCondition 根据条件获取OrderItems记录
	GetByCondition(condition map[string]interface{}) ([]*model.OrderItems, error)
	
	// FindByCondition 根据条件查找OrderItems记录 (GetByCondition 的别名)
	FindByCondition(condition map[string]interface{}) ([]*model.OrderItems, error)
	
	// SelectByCondition 根据条件选择OrderItems记录 (GetByCondition 的别名)
	SelectByCondition(condition map[string]interface{}) ([]*model.OrderItems, error)
	
	// QueryByCondition 根据条件查询OrderItems记录 (GetByCondition 的别名)
	QueryByCondition(condition map[string]interface{}) ([]*model.OrderItems, error)

	// 统计方法 - 返回数量
	// GetCount 获取OrderItems记录总数
	GetCount() (int64, error)
	
	// Count 统计OrderItems记录总数 (GetCount 的别名)
	Count() (int64, error)
	
	// CountByCondition 根据条件统计OrderItems记录数
	CountByCondition(condition map[string]interface{}) (int64, error)


	// 存在性检查方法 - 返回布尔值
	// GetExistsById 检查指定主键的OrderItems记录是否存在
	GetExistsById(id *int) (bool, error)


	// 更新方法 (UPDATE) - 返回影响行数

	// UpdateById 根据主键更新OrderItems
	UpdateById(record *model.OrderItems) (int64, error)
	
	// ModifyById 根据主键修改OrderItems (UpdateById 的别名)
	ModifyById(record *model.OrderItems) (int64, error)
	
	// EditById 根据主键编辑OrderItems (UpdateById 的别名)
	EditById(record *model.OrderItems) (int64, error)


	// UpdateByCondition 根据条件更新OrderItems记录
	UpdateByCondition(record *model.OrderItems, condition map[string]interface{}) (int64, error)

	// 删除方法 (DELETE) - 返回影响行数

	// DeleteById 根据主键删除OrderItems
	DeleteById(id *int) (int64, error)
	
	// RemoveById 根据主键移除OrderItems (DeleteById 的别名)
	RemoveById(id *int) (int64, error)
	
	// DeleteByIds 根据主键列表批量删除OrderItems
	DeleteByIds(ids []*int) (int64, error)
	
	// RemoveByIds 根据主键列表批量移除OrderItems (DeleteByIds 的别名)
	RemoveByIds(ids []*int) (int64, error)


	// DeleteByCondition 根据条件删除OrderItems记录
	DeleteByCondition(condition map[string]interface{}) (int64, error)
	
	// RemoveByCondition 根据条件移除OrderItems记录 (DeleteByCondition 的别名)
	RemoveByCondition(condition map[string]interface{}) (int64, error)


	// Example 查询方法 - 支持 GoBatis Example 功能
	// GetByExample 根据 Example 条件获取OrderItems记录
	GetByExample(example *example.Example) ([]*model.OrderItems, error)
	
	// FindByExample 根据 Example 条件查找OrderItems记录 (GetByExample 的别名)
	FindByExample(example *example.Example) ([]*model.OrderItems, error)
	
	// SelectByExample 根据 Example 条件选择OrderItems记录 (GetByExample 的别名)
	SelectByExample(example *example.Example) ([]*model.OrderItems, error)
	
	// QueryByExample 根据 Example 条件查询OrderItems记录 (GetByExample 的别名)
	QueryByExample(example *example.Example) ([]*model.OrderItems, error)
	
	// ListByExample 根据 Example 条件列出OrderItems记录 (GetByExample 的别名)
	ListByExample(example *example.Example) ([]*model.OrderItems, error)
	
	// CountByExample 根据 Example 条件统计OrderItems记录数
	CountByExample(example *example.Example) (int64, error)
	
	// UpdateByExample 根据 Example 条件更新OrderItems记录
	UpdateByExample(record *model.OrderItems, example *example.Example) (int64, error)
	
	// ModifyByExample 根据 Example 条件修改OrderItems记录 (UpdateByExample 的别名)
	ModifyByExample(record *model.OrderItems, example *example.Example) (int64, error)
	
	// EditByExample 根据 Example 条件编辑OrderItems记录 (UpdateByExample 的别名)
	EditByExample(record *model.OrderItems, example *example.Example) (int64, error)
	
	// DeleteByExample 根据 Example 条件删除OrderItems记录
	DeleteByExample(example *example.Example) (int64, error)
	
	// RemoveByExample 根据 Example 条件移除OrderItems记录 (DeleteByExample 的别名)
	RemoveByExample(example *example.Example) (int64, error)

}
