	CREATE TABLE IF NOT EXISTS users (acct varchar(30) PRIMARY KEY,pwd  varchar(20) NOT NULL,fullname  varchar(20) NOT NULL,created_at TIMESTAMP ,updated_at TIMESTAMP);
	INSERT INTO users (acct, pwd ,fullname) VALUES ('Kenny123', '0000','Kenny');
	INSERT INTO users (acct, pwd ,fullname) VALUES ('vic123', '0000','vic' );
	INSERT INTO users (acct, pwd ,fullname ,created_at, updated_at ) VALUES ('dick123', '0000' ,'dick','1999-01-08 04:05:06' , '1999-01-08 04:05:06' );
	INSERT INTO users (acct, pwd ,fullname ,created_at, updated_at ) VALUES ('liang123', '0000' ,'liang','1999-01-08 04:05:06' , '1999-01-08 04:05:06' );
