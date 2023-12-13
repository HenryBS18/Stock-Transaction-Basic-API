CREATE DATABASE stock_golang;

USE stock_golang;

INSERT INTO accounts (name) VALUES ('Henry');
INSERT INTO accounts (name) VALUES ('Andy');

INSERT INTO stocks (ticker, last_price, previous_price, open_price, volume, frequency, turnover)
VALUES
    ('BBCA', 9500, 9500, 9500, 30000000, 2230, 50300000),
    ('BMRI', 4500, 4500, 4500, 6200000, 583, 8543000);