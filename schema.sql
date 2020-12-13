CREATE TABLE todo_list
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE todo_item
(
    id           SERIAL PRIMARY KEY,
    todo_list_id INT     NOT NULL REFERENCES todo_list ON DELETE CASCADE,
    text         VARCHAR NOT NULL,
    done         BOOLEAN NOT NULL DEFAULT FALSE
);