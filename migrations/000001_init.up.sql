CREATE TABLE IF NOT EXISTS segments (
    id SERIAL PRIMARY KEY,
    slug VARCHAR UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR
);

-- many-to-many
CREATE TABLE IF NOT EXISTS users_segments (
    id SERIAL PRIMARY KEY,
    -- ttl_flag BOOLEAN,
    -- ttl_deadline TIMESTAMP
    user_id INTEGER REFERENCES users(id),
    segment_id INTEGER REFERENCES segments(id),
    UNIQUE(user_id, segment_id)
);