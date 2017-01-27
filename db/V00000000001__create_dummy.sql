create table user(
id int primary key auto_increment
,email varchar(255) not null unique
,name varchar(255) not null
,age int default 0
);

create table user_item(
id int primary key auto_increment
,user_id int not null
,name varchar(64) default ""
,foreign key (user_id) references user (id)
);
