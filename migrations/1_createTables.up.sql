# createTables: category, basic, daily_price, fail_daily_price


CREATE TABLE `category` (
    id int primary key AUTO_INCREMENT,
    name varchar(40) unique not null
);

CREATE TABLE `basic` (
    id int primary key AUTO_INCREMENT,
    category_id int,
    stock_code varchar(15) unique not null,
    stock_name varchar(40) not null,
    FOREIGN KEY(category_id) REFERENCES category(id)
);

CREATE UNIQUE INDEX `basic_index_0` ON `basic` (`stock_code`);


CREATE TABLE daily_price (
    id int primary key AUTO_INCREMENT,
    stock_code varchar(15),
    volume DECIMAL(10, 0)NOT NULL,
    opening_price DECIMAL(10, 2) NOT NULL,
    closing_price DECIMAL(10, 2) NOT NULL,
    highest_price DECIMAL(10, 2) NOT NULL,
    lowest_price DECIMAL(10, 2) NOT NULL,
    calculate_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX `daily_price_index_0` ON `daily_price` (`stock_code`, `calculate_date`);

CREATE TABLE calculate_timestamps (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `table` varchar(20),
  `calculate_timestamp` int(10)
);