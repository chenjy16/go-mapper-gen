-- 示例数据库初始化脚本
-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    avatar VARCHAR(255),
    status INTEGER DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 创建产品表
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock INTEGER DEFAULT 0,
    category_id INTEGER,
    status INTEGER DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 创建订单表
CREATE TABLE IF NOT EXISTS orders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    order_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    shipping_address TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 创建订单详情表
CREATE TABLE IF NOT EXISTS order_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- 插入示例数据
INSERT OR IGNORE INTO users (id, username, email, password, phone, status) VALUES
(1, 'admin', 'admin@example.com', 'password123', '13800138000', 1),
(2, 'john_doe', 'john@example.com', 'password456', '13800138001', 1),
(3, 'jane_smith', 'jane@example.com', 'password789', '13800138002', 1);

INSERT OR IGNORE INTO products (id, name, description, price, stock, category_id, status) VALUES
(1, 'iPhone 15', '最新款苹果手机', 7999.00, 100, 1, 1),
(2, 'MacBook Pro', '专业级笔记本电脑', 15999.00, 50, 2, 1),
(3, 'AirPods Pro', '无线降噪耳机', 1999.00, 200, 3, 1);

INSERT OR IGNORE INTO orders (id, user_id, total_amount, status, shipping_address) VALUES
(1, 1, 7999.00, 'completed', '北京市朝阳区xxx街道xxx号'),
(2, 2, 17998.00, 'pending', '上海市浦东新区xxx路xxx号'),
(3, 3, 1999.00, 'shipped', '广州市天河区xxx大道xxx号');

INSERT OR IGNORE INTO order_items (order_id, product_id, quantity, unit_price, total_price) VALUES
(1, 1, 1, 7999.00, 7999.00),
(2, 1, 1, 7999.00, 7999.00),
(2, 2, 1, 15999.00, 15999.00),
(3, 3, 1, 1999.00, 1999.00);