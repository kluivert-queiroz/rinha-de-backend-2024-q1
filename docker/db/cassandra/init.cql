CREATE KEYSPACE IF NOT EXISTS wallet WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };
USE wallet;
DROP TABLE IF EXISTS wallet.customer;
CREATE TYPE IF NOT EXISTS transaction (
      amount int,
      description text,
      type text,
      date timestamp
);
CREATE TABLE IF NOT EXISTS wallet.customer (
    id text PRIMARY KEY,
    balance int,
    "limit" int,
    transactions set<frozen<transaction>>,
    version int
);


INSERT INTO wallet.customer(id, balance, "limit", transactions, version) VALUES (  '1',  0, 100000,  {},  1 );
INSERT INTO wallet.customer (id, balance, "limit", transactions, version) VALUES (  '2',  0, 80000,  {},  1 );
INSERT INTO wallet.customer (id, balance, "limit", transactions, version) VALUES (  '3',  0, 1000000,  {},  1 );
INSERT INTO wallet.customer (id, balance, "limit", transactions, version) VALUES (  '4',  0, 10000000,  {},  1 );
INSERT INTO wallet.customer (id, balance, "limit", transactions, version) VALUES (  '5',  0, 500000,  {},  1 );