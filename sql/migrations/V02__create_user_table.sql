CREATE TABLE IF NOT EXISTS users_db.users (
   user_id SERIAL PRIMARY KEY,
   first_name VARCHAR(100) NOT NULL,
   last_name VARCHAR(100) NOT NULL
);