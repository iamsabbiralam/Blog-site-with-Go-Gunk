-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL NOT NULL,
    cat_id INTEGER NOT NULL,
    title TEXT NOT NULL UNIQUE,
    image TEXT NOT NULL,

    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
