CREATE TABLE IF NOT EXISTS Events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    date DATE,
    user_id INT
);