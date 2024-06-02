create database db_carteirago
go
use db_carteirago
go
alter login sa ENABLE
go
alter login sa with PASSWORD = 'sa@carteiraGO'

CREATE TABLE dbo.go_user(
    id_go_user bigint not null identity,
    fullname_go_user varchar(100) not null,
    email_go_user varchar(100)not null,
    password varchar(100) not null,
    socialname_go_user varchar(100) null

    PRIMARY key(id_go_user,email_go_user,fullname_go_user)    
)
go
insert into dbo.go_user values('Renan Machado','renanoliveira2199@gmail.com','carteirago123','Renan')