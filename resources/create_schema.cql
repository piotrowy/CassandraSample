CREATE KEYSPACE IF NOT EXISTS Pub
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };

USE Pub;

CREATE TABLE products (
  category varchar,
  product_name varchar,
  product_price int,
  PRIMARY KEY (category, product_name)
);

CREATE TABLE buy_counters (
  product_first varchar,
  product_second varchar,
  occurrences counter,
  PRIMARY KEY((product_first, product_second))
);

CREATE TABLE history (
  user_id varchar,
  product_name varchar,
  product_price int,
  date timestamp,
  PRIMARY KEY (user_id, date)
);

CREATE TABLE recommendation (
  product_name varchar,
  other_product_name varchar,
  other_product_price int,
  PRIMARY KEY (product_name, other_product_name)
);