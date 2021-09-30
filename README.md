# smrt-alive-app

## How to start a Project

go mod init {url of the github project}

## SMRT alive app
https://talks.golang.org/2014/names.slide#1

## Run 
go run cmd/web/* -addr=":{insert address number here}"

## MySQL
This webiste is build in mysql 
to access type sudo mysql to access the terminal
userbox is the the database name

-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE userbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

The cmd director will contain the application-specific code
pkg will contain the ancillary non-application - specific code e.g. validation helpers and sql
ui directory will contain the user interface assets used by the web application 

CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT ON userbox.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'whaleredchurch';

mysql -D userbox -u web -p

The table we used to create

CREATE TABLE students (
    `forename` VARCHAR(100) NOT NULL,
    `surname`  VARCHAR(100) NOT NULL,
    `email` VARCHAR(200) NOT NULL,
    `password` TEXT NOT NULL,
    `yofe` INT NOT NULL,
    `tutor` VARCHAR(100) NOT NULL,
    `aliveorganizationandtime` VARCHAR(100),
    `aliveorganizationandtimeev` TEXT,
     PRIMARY KEY (`email`)
);

## TO access sql 
sudo mysql
USE userbox


https://wtools.io/generate-sql-create-table
