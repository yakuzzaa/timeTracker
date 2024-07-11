-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP,
    total INTERVAL
);

-- +goose StatementBegin
-- Function to set end_time and calculate total
CREATE OR REPLACE FUNCTION set_end_time_and_total() RETURNS TRIGGER AS $set_end_time_and_total$
BEGIN
    NEW.end_time = CURRENT_TIMESTAMP;
    NEW.total = EXTRACT(EPOCH FROM (NEW.end_time - NEW.start_time)) / 60;
    RETURN NEW;
END;
$set_end_time_and_total$ LANGUAGE plpgsql;


-- Trigger to set end_time and calculate total on update
CREATE TRIGGER set_end_time_and_total_trigger
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    WHEN (OLD.end_time IS NULL)
    EXECUTE FUNCTION set_end_time_and_total();
-- +goose StatementEnd

-- +goose Down
DROP TABLE tasks;

-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_end_time_and_total_trigger ON tasks;
DROP FUNCTION IF EXISTS set_end_time_and_total;
-- +goose StatementEnd

