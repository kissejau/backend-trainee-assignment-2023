CREATE TABLE IF NOT EXISTS segments (
    id SERIAL PRIMARY KEY,
    slug VARCHAR UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    slug VARCHAR,
    type VARCHAR,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- insert data to logs table on new record in users_segments
CREATE OR REPLACE FUNCTION users_segments_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO logs (user_id, slug, type, created_at)
    VALUES (NEW.user_id, (SELECT slug FROM segments WHERE id = NEW.segment_id), 'Add', NOW());
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- insert data to logs table on deleting record from users_segments 
CREATE OR REPLACE FUNCTION users_segments_delete_trigger()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO logs (user_id, slug, type, created_at)
    VALUES (OLD.user_id, (SELECT slug FROM segments WHERE id = OLD.segment_id), 'Delete', NOW());
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- many-to-many
CREATE TABLE IF NOT EXISTS users_segments (
    id SERIAL PRIMARY KEY,
    deadline TIMESTAMP,
    user_id INTEGER REFERENCES users(id),
    segment_id INTEGER REFERENCES segments(id),
    UNIQUE(user_id, segment_id)
);

CREATE TRIGGER users_segments_after_insert
AFTER INSERT ON users_segments
FOR EACH ROW
EXECUTE FUNCTION users_segments_insert_trigger();

CREATE TRIGGER users_segments_after_delete
AFTER DELETE ON users_segments
FOR EACH ROW
EXECUTE FUNCTION users_segments_delete_trigger();