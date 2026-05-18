-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    id               uuid        PRIMARY KEY DEFAULT gen_random_uuid(),
    title            text        NOT NULL,
    slug             text        UNIQUE NOT NULL,
    markdown_content text        NOT NULL,
    meta_description text        NULL,
    cover_image_url  text        NULL,
    youtube_links    jsonb       NULL,
    published_at     timestamptz NULL,
    user_id          uuid        NOT NULL,
    created_at       timestamptz NOT NULL DEFAULT now(),
    updated_at       timestamptz NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS posts;
