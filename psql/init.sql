	CREATE TABLE IF NOT EXISTS users (acct varchar(30) PRIMARY KEY,pwd  varchar(20) NOT NULL ,created_at TIMESTAMP ,updated_at TIMESTAMP);
	INSERT INTO users (acct, pwd ) VALUES ('Kenny', '0000' );
	INSERT INTO users (acct, pwd ) VALUES ('vic', '0000' );
	INSERT INTO users (acct, pwd ,created_at, updated_at ) VALUES ('dick', '0000' , '1999-01-08 04:05:06' , '1999-01-08 04:05:06' );
	SELECT * FROM users ;