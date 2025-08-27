-- =================================
-- MICROSERVICES DATABASE SETUP
-- =================================

-- Criar databases para cada microservice
CREATE DATABASE IF NOT EXISTS `order`;
CREATE DATABASE IF NOT EXISTS `payment`;
CREATE DATABASE IF NOT EXISTS `shipping`;

-- =================================
-- ORDER SERVICE DATABASE
-- =================================
USE `order`;

CREATE TABLE orders (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    customer_id BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    total_amount DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE order_items (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    product_code VARCHAR(100) NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    subtotal DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    INDEX idx_order_id (order_id),
    INDEX idx_product_code (product_code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Tabela de produtos para validação e controle de estoque
CREATE TABLE products (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    product_code VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INT NOT NULL DEFAULT 0,
    min_stock INT NOT NULL DEFAULT 0,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_product_code (product_code),
    INDEX idx_active (active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- =================================
-- PAYMENT SERVICE DATABASE  
-- =================================
USE `payment`;

CREATE TABLE payments (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    customer_id BIGINT NOT NULL,
    order_id BIGINT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    payment_method VARCHAR(50) DEFAULT 'CREDIT_CARD',
    transaction_id VARCHAR(255),
    processed_at TIMESTAMP NULL,
    created_at BIGINT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_order_id (order_id),
    INDEX idx_customer_id (customer_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- =================================
-- SHIPPING SERVICE DATABASE
-- =================================
USE `shipping`;

CREATE TABLE shippings (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    delivery_days INT NOT NULL,
    estimated_delivery DATE,
    tracking_code VARCHAR(100),
    status VARCHAR(50) NOT NULL DEFAULT 'PROCESSING',
    total_items INT NOT NULL DEFAULT 0,
    shipping_cost DECIMAL(10,2) DEFAULT 0.00,
    created_at BIGINT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_order_id (order_id),
    INDEX idx_status (status),
    INDEX idx_tracking_code (tracking_code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE shipping_items (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    shipping_id BIGINT NOT NULL,
    product_code VARCHAR(100) NOT NULL,
    quantity INT NOT NULL,
    weight DECIMAL(8,2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (shipping_id) REFERENCES shippings(id) ON DELETE CASCADE,
    INDEX idx_shipping_id (shipping_id),
    INDEX idx_product_code (product_code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- =================================
-- DADOS INICIAIS
-- =================================
USE `order`;

-- Inserir produtos de exemplo
INSERT INTO products (product_code, name, description, price, stock_quantity, min_stock) VALUES
('PROD001', 'Produto Premium 1', 'Produto de alta qualidade com garantia estendida', 10.50, 100, 10),
('PROD002', 'Produto Especial 2', 'Item exclusivo com características únicas', 25.00, 50, 5),
('PROD003', 'Produto Padrão 3', 'Produto básico com boa relação custo-benefício', 15.75, 75, 8),
('PROD004', 'Produto Deluxe 4', 'Versão premium com acabamento refinado', 30.00, 25, 3),
('PROD005', 'Produto Econômico 5', 'Opção acessível mantendo a qualidade', 12.99, 200, 20);

-- =================================
-- ÍNDICES ADICIONAIS PARA PERFORMANCE
-- =================================

-- Order service indexes
USE `order`;
ALTER TABLE orders ADD INDEX idx_customer_status (customer_id, status);
ALTER TABLE orders ADD INDEX idx_created_at (created_at);

-- Payment service indexes  
USE `payment`;
ALTER TABLE payments ADD INDEX idx_created_status (created_at, status);

-- Shipping service indexes
USE `shipping`;
ALTER TABLE shippings ADD INDEX idx_created_delivery (created_at, delivery_days);

-- =================================
-- VIEWS ÚTEIS PARA RELATÓRIOS
-- =================================

USE `order`;
CREATE VIEW order_summary AS
SELECT 
    o.id,
    o.customer_id,
    o.status,
    o.total_amount,
    COUNT(oi.id) as total_items,
    SUM(oi.quantity) as total_quantity,
    o.created_at
FROM orders o
LEFT JOIN order_items oi ON o.id = oi.order_id AND oi.deleted_at IS NULL
WHERE o.deleted_at IS NULL
GROUP BY o.id;

-- =================================
-- CONFIGURAÇÕES DE SEGURANÇA
-- =================================

-- Criar usuário específico para a aplicação (opcional)
-- CREATE USER 'microservices'@'%' IDENTIFIED BY 'app_password';
-- GRANT SELECT, INSERT, UPDATE, DELETE ON `order`.* TO 'microservices'@'%';
-- GRANT SELECT, INSERT, UPDATE, DELETE ON `payment`.* TO 'microservices'@'%';
-- GRANT SELECT, INSERT, UPDATE, DELETE ON `shipping`.* TO 'microservices'@'%';
-- FLUSH PRIVILEGES;
