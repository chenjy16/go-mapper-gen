-- Users 相关 SQL 语句
-- 表名: users

-- 查询所有记录
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at
FROM users;


-- 根据主键查询
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at
FROM users
WHERE id = ?;


-- 插入记录
INSERT INTO users (
    username,
    email,
    password,
    phone,
    avatar,
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
INSERT INTO users (
    username,
    email,
    password,
    phone,
    avatar,
    status,
    created_at,
    updated_at
) VALUES 
(?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?),
(?, ?, ?, ?, ?, ?, ?, ?);


-- 根据主键更新
UPDATE users
SET username = ?,
    email = ?,
    password = ?,
    phone = ?,
    avatar = ?,
    status = ?,
    created_at = ?,
    updated_at = ?
WHERE id = ?;

-- 根据主键删除
DELETE FROM users
WHERE id = ?;

-- 批量删除
DELETE FROM users
WHERE id IN (?, ?, ?);


-- 分页查询
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at
FROM users
ORDER BY id
LIMIT ? OFFSET ?;

-- 统计总数
SELECT COUNT(*) FROM users;


-- 检查记录是否存在
SELECT COUNT(*) FROM users
WHERE id = ?;


-- 条件查询示例
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at
FROM users
WHERE 1=1
  -- AND username = ?
  -- AND email = ?
  -- AND password = ?
  -- AND phone = ?
  -- AND avatar = ?
  -- AND status = ?
  -- AND created_at = ?
  -- AND updated_at = ?
ORDER BY id;

-- 模糊查询示例
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at
FROM users
WHERE 1=1
  -- AND username LIKE CONCAT('%', ?, '%')
  -- AND email LIKE CONCAT('%', ?, '%')
  -- AND password LIKE CONCAT('%', ?, '%')
  -- AND phone LIKE CONCAT('%', ?, '%')
  -- AND avatar LIKE CONCAT('%', ?, '%')
ORDER BY id;

-- 范围查询示例
SELECT id, username, email, password, phone, avatar, status, created_at, updated_at
FROM users
WHERE 1=1
  -- AND id BETWEEN ? AND ?
  -- AND status BETWEEN ? AND ?
  -- AND created_at BETWEEN ? AND ?
  -- AND updated_at BETWEEN ? AND ?
ORDER BY id;
