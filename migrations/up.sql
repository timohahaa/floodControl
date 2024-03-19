CREATE TABLE requests (
    req_id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    req_timestamp TIMESTAMP NOT NULL DEFAULT NOW()
);
