CREATE TABLE master.dbo.book (
	id varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	name varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	author varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	CONSTRAINT PK__book__3213E83F1A4FA0E0 PRIMARY KEY (id)
) GO;


INSERT INTO master.dbo.book (id,name,author) VALUES 
('045ce937-2122-4804-8d05-e30f6e07acfc','book003','A.B')
,('1958fd68-422f-435e-ab73-81e3eb472799','book001','A.U')
,('3527285a-64e7-46b3-95d6-2b9d849b6504','Book002','B.C')
;