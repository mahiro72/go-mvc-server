-- migrate:up
ALTER TABLE users
ADD mail varchar(255);

-- migrate:down
ALTER TABLE users
DROP mail;
