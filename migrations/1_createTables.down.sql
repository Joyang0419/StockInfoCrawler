# dropTables: category, basic, daily_price, fail_daily_price

ALTER TABLE basic DROP FOREIGN KEY basic_ibfk_1;
DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS basic;
DROP TABLE IF EXISTS daily_price;
DROP TABLE IF EXISTS fail_daily_price;
DROP TABLE IF EXISTS calculate_timestamps;