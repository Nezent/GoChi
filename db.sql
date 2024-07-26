-- DELETE FROM account
-- WHERE account_id > 2;

-- SELECT * FROM account;

-- SELECT * FROM transaction;

-- CREATE TABLE transaction (
--     transaction_id SERIAL PRIMARY KEY,
--     account_id SERIAL REFERENCES account(account_id),
--     transaction_type VARCHAR(20),
--     amount NUMERIC,
--     transaction_date TIMESTAMP
-- );

-- DROP TABLE transaction;