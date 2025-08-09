-- Orders 相关 SQL 语句
-- 表名: orders

-- 查询所有记录
SELECT id, order_no, user_id, total_amount, status, payment_method, shipping_address, remark, created_at, updated_at
FROM orders;


-- 根据主键查询
SELECT id, order_no, user_id, total_amount, status, payment_method, shipping_address, remark, created_at, updated_at
FROM orders
WHERE id = ?;


-- 插入记录
INSERT INTO orders (
    order_no,
    user_id,
    total_amount,
    status,
    payment_method,
    shipping_address,
    remark,
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
    ?,
    ?
);

-- 批量插入记录
INSERT INTO orders (
    order_no,
    user_id,
    total_amount,
    status,
    payment_method,
    shipping_address,
    remark,
    created_at,
    updated_at
) VALUES 
(?, ?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?, ?);


-- 根据主键更新
UPDATE orders
SET order_no = ?,
    user_id = ?,
    total_amount = ?,
    status = ?,
    payment_method = ?,
    shipping_address = ?,
    remark = ?,
    created_at = ?,
    updated_at = ?
WHERE id = ?;

-- 根据主键删除
DELETE FROM orders
WHERE id = ?;

-- 批量删除
DELETE FROM orders
WHERE id IN (?, ?, ?);


-- 分页查询
SELECT id, order_no, user_id, total_amount, status, payment_method, shipping_address, remark, created_at, updated_at
FROM orders
ORDER BY id
LIMIT ? OFFSET ?;

-- 统计总数
SELECT COUNT(*) FROM orders;


-- 检查记录是否存在
SELECT COUNT(*) FROM orders
WHERE id = ?;


-- 条件查询示例
SELECT id, order_no, user_id, total_amount, status, payment_method, shipping_address, remark, created_at, updated_at
FROM orders
WHERE 1=1
  -- AND order_no = ?
  -- AND user_id = ?
  -- AND total_amount = ?
  -- AND status = ?
  -- AND payment_method = ?
  -- AND shipping_address = ?
  -- AND remark = ?
  -- AND created_at = ?
  -- AND updated_at = ?
ORDER BY id;

-- 模糊查询示例
SELECT id, order_no, user_id, total_amount, status, payment_method, shipping_address, remark, created_at, updated_at
FROM orders
WHERE 1=1
  -- AND order_no LIKE CONCAT('%', ?, '%')
  -- AND shipping_address LIKE CONCAT('%', ?, '%')
  -- AND remark LIKE CONCAT('%', ?, '%')
ORDER BY id;

-- 范围查询示例
SELECT id, order_no, user_id, total_amount, status, payment_method, shipping_address, remark, created_at, updated_at
FROM orders
WHERE 1=1
  -- AND id BETWEEN ? AND ?
  -- AND user_id BETWEEN ? AND ?
  -- AND status BETWEEN ? AND ?
  -- AND payment_method BETWEEN ? AND ?
  -- AND created_at BETWEEN ? AND ?
  -- AND updated_at BETWEEN ? AND ?
ORDER BY id;
