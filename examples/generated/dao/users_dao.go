package dao

import (
	model "go-mapper-gen/examples/generated/model"
	"gobatis/core/example"
)

// UsersDAO Users 数据访问接口
// 严格遵循 GoBatis 框架方法命名规则和返回值规范
type UsersDAO interface {
	// 插入方法 (INSERT) - 返回影响行数
	// Insert 插入单个Users记录
	Insert(record *model.Users) (int64, error)
	
	// InsertBatch 批量插入Users记录
	InsertBatch(records []*model.Users) (int64, error)
	
	// Add 添加Users记录 (Insert 的别名)
	Add(record *model.Users) (int64, error)
	
	// Create 创建Users记录 (Insert 的别名)
	Create(record *model.Users) (int64, error)
	
	// Save 保存Users记录 (Insert 的别名)
	Save(record *model.Users) (int64, error)


	// 查询方法 (SELECT) - 返回查询结果
	// GetById 根据主键获取Users
	GetById(id *int) (*model.Users, error)
	
	// FindById 根据主键查找Users (GetById 的别名)
	FindById(id *int) (*model.Users, error)
	
	// SelectById 根据主键选择Users (GetById 的别名)
	SelectById(id *int) (*model.Users, error)


	// GetAll 获取所有Users记录
	GetAll() ([]*model.Users, error)
	
	// FindAll 查找所有Users记录 (GetAll 的别名)
	FindAll() ([]*model.Users, error)
	
	// SelectAll 选择所有Users记录 (GetAll 的别名)
	SelectAll() ([]*model.Users, error)
	
	// ListAll 列出所有Users记录 (GetAll 的别名)
	ListAll() ([]*model.Users, error)
	
	// QueryAll 查询所有Users记录 (GetAll 的别名)
	QueryAll() ([]*model.Users, error)
	
	// GetByPage 分页获取Users记录
	GetByPage(offset, limit int) ([]*model.Users, error)
	
	// FindByPage 分页查找Users记录 (GetByPage 的别名)
	FindByPage(offset, limit int) ([]*model.Users, error)
	
	// SelectByPage 分页选择Users记录 (GetByPage 的别名)
	SelectByPage(offset, limit int) ([]*model.Users, error)
	
	// GetByCondition 根据条件获取Users记录
	GetByCondition(condition map[string]interface{}) ([]*model.Users, error)
	
	// FindByCondition 根据条件查找Users记录 (GetByCondition 的别名)
	FindByCondition(condition map[string]interface{}) ([]*model.Users, error)
	
	// SelectByCondition 根据条件选择Users记录 (GetByCondition 的别名)
	SelectByCondition(condition map[string]interface{}) ([]*model.Users, error)
	
	// QueryByCondition 根据条件查询Users记录 (GetByCondition 的别名)
	QueryByCondition(condition map[string]interface{}) ([]*model.Users, error)

	// 统计方法 - 返回数量
	// GetCount 获取Users记录总数
	GetCount() (int64, error)
	
	// Count 统计Users记录总数 (GetCount 的别名)
	Count() (int64, error)
	
	// CountByCondition 根据条件统计Users记录数
	CountByCondition(condition map[string]interface{}) (int64, error)


	// 存在性检查方法 - 返回布尔值
	// GetExistsById 检查指定主键的Users记录是否存在
	GetExistsById(id *int) (bool, error)


	// 更新方法 (UPDATE) - 返回影响行数

	// UpdateById 根据主键更新Users
	UpdateById(record *model.Users) (int64, error)
	
	// ModifyById 根据主键修改Users (UpdateById 的别名)
	ModifyById(record *model.Users) (int64, error)
	
	// EditById 根据主键编辑Users (UpdateById 的别名)
	EditById(record *model.Users) (int64, error)


	// UpdateByCondition 根据条件更新Users记录
	UpdateByCondition(record *model.Users, condition map[string]interface{}) (int64, error)

	// 删除方法 (DELETE) - 返回影响行数

	// DeleteById 根据主键删除Users
	DeleteById(id *int) (int64, error)
	
	// RemoveById 根据主键移除Users (DeleteById 的别名)
	RemoveById(id *int) (int64, error)
	
	// DeleteByIds 根据主键列表批量删除Users
	DeleteByIds(ids []*int) (int64, error)
	
	// RemoveByIds 根据主键列表批量移除Users (DeleteByIds 的别名)
	RemoveByIds(ids []*int) (int64, error)


	// DeleteByCondition 根据条件删除Users记录
	DeleteByCondition(condition map[string]interface{}) (int64, error)
	
	// RemoveByCondition 根据条件移除Users记录 (DeleteByCondition 的别名)
	RemoveByCondition(condition map[string]interface{}) (int64, error)


	// Example 查询方法 - 支持 GoBatis Example 功能
	// GetByExample 根据 Example 条件获取Users记录
	GetByExample(example *example.Example) ([]*model.Users, error)
	
	// FindByExample 根据 Example 条件查找Users记录 (GetByExample 的别名)
	FindByExample(example *example.Example) ([]*model.Users, error)
	
	// SelectByExample 根据 Example 条件选择Users记录 (GetByExample 的别名)
	SelectByExample(example *example.Example) ([]*model.Users, error)
	
	// QueryByExample 根据 Example 条件查询Users记录 (GetByExample 的别名)
	QueryByExample(example *example.Example) ([]*model.Users, error)
	
	// ListByExample 根据 Example 条件列出Users记录 (GetByExample 的别名)
	ListByExample(example *example.Example) ([]*model.Users, error)
	
	// CountByExample 根据 Example 条件统计Users记录数
	CountByExample(example *example.Example) (int64, error)
	
	// UpdateByExample 根据 Example 条件更新Users记录
	UpdateByExample(record *model.Users, example *example.Example) (int64, error)
	
	// ModifyByExample 根据 Example 条件修改Users记录 (UpdateByExample 的别名)
	ModifyByExample(record *model.Users, example *example.Example) (int64, error)
	
	// EditByExample 根据 Example 条件编辑Users记录 (UpdateByExample 的别名)
	EditByExample(record *model.Users, example *example.Example) (int64, error)
	
	// DeleteByExample 根据 Example 条件删除Users记录
	DeleteByExample(example *example.Example) (int64, error)
	
	// RemoveByExample 根据 Example 条件移除Users记录 (DeleteByExample 的别名)
	RemoveByExample(example *example.Example) (int64, error)

}
