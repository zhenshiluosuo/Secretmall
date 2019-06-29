create table user(
  id int primary key ,
  name varchar(20) not null ,
  password varchar(20) not null,
  permission enum('0','1') default 1 not null
);

insert into user values(1,'gofer','gofer','0');
insert into user values(2,'your','father','1');

create table item(
  id int primary key ,
  name varchar(30) not null ,
  price float not null ,
  amount  int not null ,
  descirption varchar(50)
);

insert into item values(1,'useless toothbrush',12.0,2,'');

create table cert(
  id  int primary key ,
  user_id int not null ,
  item_id int not null ,
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (item_id) REFERENCES item(id)
)