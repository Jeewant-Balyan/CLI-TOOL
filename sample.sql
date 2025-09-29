CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE
);

CREATE TABLE orders (
  id INTEGER PRIMARY KEY,
  user_id INTEGER,
  total INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO users (name,email) VALUES ('Alice','alice@example.com'), ('Bob','bob@example.com');
INSERT INTO orders (user_id,total) VALUES (1,100), (2,200);
