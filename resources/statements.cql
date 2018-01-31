INSERT INTO products (category, product_name, product_price) VALUES (?, ?, ?);
UPDATE buy_counters set occurrences = occurrences + 1 where product_first = ? and product_second = ?;
INSERT INTO history (user_id, product_name, product_price, date) VALUES (?, ?, ?, ?);
INSERT INTO recommendation (product_name, other_product_name, other_product_price) VALUES (?, ?, ?);
SELECT occurrences FROM buy_counters WHERE product_first = ? and product_second = ?;
SELECT * FROM recommendation WHERE product_name = ?;
DELETE FROM recommendation WHERE product_name = ? and other_product_name = ?;
SELECT * FROM history WHERE user_id = ?;
SELECT * FROM products WHERE category = ?;