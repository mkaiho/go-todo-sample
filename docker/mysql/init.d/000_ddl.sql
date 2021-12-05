CREATE TABLE users (
  id varchar(40) NOT NULL,
  name varchar(40) NOT NULL,
  email varchar(50) NOT NULL,
  password varchar(50) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY email (email)
);