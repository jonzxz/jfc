CREATE USER 'bizuser'@'%' IDENTIFIED BY 'secretpassword';
GRANT ALL PRIVILEGES ON JFC.* to 'bizuser'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

CREATE DATABASE JFC;

USE JFC;

DROP TABLE IF EXISTS PERSON;
CREATE TABLE PERSON (
  ID INT AUTO_INCREMENT NOT NULL,
  NAME VARCHAR(255) NOT NULL,
  TELEGRAM_ID VARCHAR(255) NOT NULL,
  HOUSEHOLD VARCHAR(255) NOT NULL,
  CREATED_TS BIGINT NOT NULL,
  LAST_UPDATED_TS BIGINT,
  PRIMARY KEY (`ID`)
);

DROP TABLE IF EXISTS PAYMENT;
CREATE TABLE PAYMENT (
  ID INT AUTO_INCREMENT NOT NULL,
  TYPE VARCHAR(255) NOT NULL,
  REMARKS VARCHAR(255) NOT NULL,
  TOTAL_AMOUNT FLOAT(2) NOT NULL,
  HOUSEHOLD VARCHAR(255) NOT NULL,
  CREATED_TS BIGINT NOT NULL,
  LAST_UPDATED_TS BIGINT,
  PRIMARY KEY (`ID`)
);

DROP TABLE IF EXISTS PAYMENT_DUE;
CREATE TABLE PAYMENT_DUE (
  ID INT AUTO_INCREMENT NOT NULL,
  PAYER_ID INT NOT NULL,
  PAYMENT_ID INT NOT NULL,
  PAYABLE_AMOUNT FLOAT(2) NOT NULL,
  PAYMENT_DUE_TS BIGINT,
  PAID BOOL NOT NULL,
  CREATED_TS BIGINT NOT NULL,
  LAST_UPDATED_TS BIGINT,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (PAYER_ID) REFERENCES PERSON(ID),
  FOREIGN KEY (PAYMENT_ID) REFERENCES PAYMENT(ID)
);

INSERT INTO PERSON (NAME, TELEGRAM_ID, HOUSEHOLD, CREATED_TS) VALUES (
  "Jonathan", "@jonny", "123", UNIX_TIMESTAMP()
);

INSERT INTO PAYMENT (CREATED_TS, TYPE, REMARKS, TOTAL_AMOUNT, HOUSEHOLD) VALUES (
  UNIX_TIMESTAMP(), "Utilities", "Payment for SP in Nov", 288.32, "123"
);

INSERT INTO PAYMENT (CREATED_TS, TYPE, REMARKS, TOTAL_AMOUNT, HOUSEHOLD) VALUES (
  UNIX_TIMEStAMP(), "Utilities", "Payment for instalment", 244.11, "B456"
);


-- INSERT INTO PAYMENT_DUE (PAYER_ID, PAYMENT_DUE_TIMESTAMP, PAYABLE_AMOUNT, PAID, PAYMENT_ID) VALUES (
  -- (SELECT ID FROM PERSON WHERE NAME='Jonathan'), 1703991712, 90.28, FALSE, (SELECT ID FROM PAYMENT WHERE ID=1)
-- );