-- OrderItems 相关 SQL 语句
-- 表名: order_items

-- 查询所有记录
SELECT id, order_id, product_id, quantity, unit_price, total_price, created_at
FROM order_items;


-- 根据主键查询
SELECT id, order_id, product_id, quantity, unit_price, total_price, created_at
FROM order_items
WHERE id = ?;


-- 插入记录
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    unit_price,
    total_price,
    created_at
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);

-- 批量插入记录
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    unit_price,
    total_price,
    created_at
) VALUES 
(?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?);


-- 根据主键更新
UPDATE order_items
SET order_id = ?,
    product_id = ?,
    quantity = ?,
    unit_price = ?,
    total_price = ?,
    created_at = ?
WHERE id = ?;

-- 根据主键删除
DELETE FROM order_items
WHERE id = ?;

-- 批量删除
DELETE FROM order_items
WHERE id IN (?, ?, ?);


-- 分页查询
SELECT id, order_id, product_id, quantity, unit_price, total_price, created_at
FROM order_items
ORDER BY id
LIMIT ? OFFSET ?;

-- 统计总数
SELECT COUNT(*) FROM order_items;


-- 检查记录是否存在
SELECT COUNT(*) FROM order_items
WHERE id = ?;


-- 条件查询示例
SELECT id, order_id, product_id, quantity, unit_price, total_price, created_at
FROM order_items
WHERE 1=1
  -- AND order_id = ?
  -- AND product_id = ?
  -- AND quantity = ?
  -- AND unit_price = ?
  -- AND total_price = ?
  -- AND created_at = ?
ORDER BY id;

-- 模糊查询示例
SELECT id, order_id, product_id, quantity, unit_price, total_price, created_at
FROM order_items
WHERE 1=1
ORDER BY id;

-- 范围查询示例
SELECT id, order_id, product_id, quantity, unit_price, total_price, created_at
FROM order_items
WHERE 1=1
  -- AND id BETWEEN ? AND ?
  -- AND order_id BETWEEN ? AND ?
  -- AND product_id BETWEEN ? AND ?
  -- AND quantity BETWEEN ? AND ?
  -- AND unit_price BETWEEN ? AND ?
  -- AND total_price BETWEEN ? AND ?
  -- AND created_at BETWEEN ? AND ?
ORDER BY id;
