CREATE TABLE users
(
    id       VARCHAR(36) PRIMARY KEY,
    username VARCHAR(36) NOT NULL,
    email    VARCHAR(36) NOT NULL,
    password VARCHAR(36) NOT NULL,
    status   VARCHAR(15) NOT NULL
);