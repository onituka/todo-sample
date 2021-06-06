DROP DATABASE IF EXISTS sample_db;
CREATE DATABASE sample_db;
USE sample_db;

CREATE TABLE priorities (
id INT NOT NULL AUTO_INCREMENT,
priority VARCHAR(10) NOT NULL,
color CHAR(7) NOT NULL,
PRIMARY KEY(id)
);

CREATE TABLE todos (
id INT NOT NULL AUTO_INCREMENT,
title VARCHAR(40) NOT NULL,
memo VARCHAR(250) DEFAULT NULL,
implementation_date DATE NOT NULL,
due_date DATE NOT NULL,
priorities_id INT DEFAULT NULL,
complete_flag BIT(1) NOT NULL DEFAULT b'0',
PRIMARY KEY(id),
FOREIGN KEY fk_priorities_id(priorities_id)
  REFERENCES priorities(id)
  ON DELETE SET NULL ON UPDATE CASCADE
);
