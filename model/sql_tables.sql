CREATE TABLE users (
    user_id INTEGER NOT NULL,
    user_name CHAR(10) NOT NULL UNIQUE,
    user_email CHAR(25) NOT NULL UNIQUE,
    user_pass PASSWORD NOT NULL,
    user_type TEXT NOT NULL DEFAULT member,
    time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("user_id" AUTOINCREMENT)
);

CREATE TABLE posts (
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    img_url TEXT,
    body TEXT NOT NULL,
    time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("post_id" AUTOINCREMENT),
    FOREIGN KEY("user_id") REFERENCES users("user_id")
);

CREATE TABLE interaction_posts (
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    interaction TEXT NOT NULL,
    FOREIGN KEY("post_id") REFERENCES posts("post_id"),
    FOREIGN KEY("user_id") REFERENCES users("user_id")
);

CREATE TABLE comments (
    comment_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("comment_id" AUTOINCREMENT),
    FOREIGN KEY("post_id") REFERENCES posts("post_id"),
    FOREIGN KEY("user_id") REFERENCES users("user_id")
);

CREATE TABLE interaction_comments (
    comment_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    interaction TEXT NOT NULL,
    FOREIGN KEY("comment_id") REFERENCES comments("comment_id"),
    FOREIGN KEY("post_id") REFERENCES posts("post_id"),
    FOREIGN KEY("user_id") REFERENCES users("user_id")
);

CREATE TABLE requests (
    request_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("request_id" AUTOINCREMENT),
    FOREIGN KEY("user_id") REFERENCES users("user_id")
);

CREATE TABLE actions (
    request_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY("request_id") REFERENCES requests("request_id"),
    FOREIGN KEY("user_id") REFERENCES users("user_id")
);

/*
sqlite3
sqlite3 mydb.db << create db
.open mydb.db
add table
.quit
*/

INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)
SELECT user_id, user_name FROM users WHERE user_id = 1
SELECT user_pass FROM users WHERE user_email = 1