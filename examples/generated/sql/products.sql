-- Products 相关 SQL 语句
-- 表名: products

-- 查询所有记录
SELECT id, name, description, price, stock, category_id, status, created_at, updated_at
FROM products;


-- 根据主键查询
SELECT id, name, description, price, stock, category_id, status, created_at, updated_at
FROM products
WHERE id = ?;


-- 插入记录
INSERT INTO products (
    name,
    description,
    price,
    stock,
    category_id,
    status,
    created_at,
    updated_at
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);

-- 批量插入记录
INSERT INTO products (
    name,
    description,
    price,
    stock,
    category_id,
    status,
    created_at,
    updated_at
) VALUES 
(?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?);


-- 根据主键更新
UPDATE products
SET name = ?,
    description = ?,
    price = ?,
    stock = ?,
    category_id = ?,
    status = ?,
    created_at = ?,
    updated_at = ?
WHERE id = ?;

-- 根据主键删除
DELETE FROM products
WHERE id = ?;

-- 批量删除
DELETE FROM products
WHERE id IN (?, ?, ?);


-- 分页查询
SELECT id, name, description, price, stock, category_id, status, created_at, updated_at
FROM products
ORDER BY id
LIMIT ? OFFSET ?;

-- 统计总数
SELECT COUNT(*) FROM products;


-- 检查记录是否存在
SELECT COUNT(*) FROM products
WHERE id = ?;


-- 条件查询示例
SELECT id, name, description, price, stock, category_id, status, created_at, updated_at
FROM products
WHERE 1=1
  -- AND name = ?
  -- AND description = ?
  -- AND price = ?
  -- AND stock = ?
  -- AND category_id = ?
  -- AND status = ?
  -- AND created_at = ?
  -- AND updated_at = ?
ORDER BY id;

-- 模糊查询示例
SELECT id, name, description, price, stock, category_id, status, created_at, updated_at
FROM products
WHERE 1=1
  -- AND name LIKE CONCAT('%', ?, '%')
  -- AND description LIKE CONCAT('%', ?, '%')
ORDER BY id;

-- 范围查询示例
SELECT id, name, description, price, stock, category_id, status, created_at, updated_at
FROM products
WHERE 1=1
  -- AND id BETWEEN ? AND ?
  -- AND price BETWEEN ? AND ?
  -- AND stock BETWEEN ? AND ?
  -- AND category_id BETWEEN ? AND ?
  -- AND status BETWEEN ? AND ?
  -- AND created_at BETWEEN ? AND ?
  -- AND updated_at BETWEEN ? AND ?
ORDER BY id;
