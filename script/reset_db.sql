DROP DATABASE If EXISTS crudapp;
CREATE DATABASE crudapp;

use crudapp;

CREATE TABLE `users`(
  `userid`		int(10) AUTO_INCREMENT,
  `userpass`	varchar(10) DEFAULT NULL,
  `username`	varchar(10) DEFAULT NULL,

	PRIMARY KEY(`userid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `foods`(
  `shopid`     	int(10)		NOT NULL,
  `foodid`     	int(10)		NOT NULL,
  `foodname`   	varchar(255) DEFAULT NULL,
  `foodsize`    varchar(255) DEFAULT NULL,
  `foodsummary`	varchar(255) DEFAULT NULL,
  `foodgenre`	varchar(255) DEFAULT NULL,
  `foodimage`	varchar(255) DEFAULT NULL,
  `foodprice`	varchar(255) DEFAULT NULL,
  
  PRIMARY KEY(shopid,foodid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into foods (shopid,foodid,foodname,foodsize,foodsummary,foodgenre,foodimage,foodprice) values ('1','1','らうめん','大盛り','んまい','ラーメン','https://ichiran.com/img/page/menu-menu_main01.png','780円');
​
CREATE TABLE `shops`(
  `shopid`     int(10) NOT NULL,
  `shopname`   varchar(255) DEFAULT NULL,
  `shopsummary`     varchar(255) DEFAULT NULL,
  `shoptel`     varchar(255) DEFAULT NULL,
  `shoptable`     varchar(255) DEFAULT NULL,
  `shopaddress`     varchar(255) DEFAULT NULL,
  `shopgenre`     varchar(255) DEFAULT NULL,  
  PRIMARY KEY(shopid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into shops (shopid,shopname,shopsummary,shoptel,shoptable,shopaddress,shopgenre) values('1','んまいラーメン','らうめんがんまい','0120-919-919','8','兵庫県神戸市中央区北野坂1-1','ラーメン');
​
CREATE TABLE `cashiers`(
  `cashid`     int(10)	NOT NULL,
  `shopid`   int(10)	NOT NULL,
  `tableid`     int(10) NOT NULL,
  `foodid`     int(10)	NOT NULL,
  PRIMARY KEY(cashid,shopid,tableid,foodid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
​
CREATE TABLE `tables`(
  `tableid`     int(10) NOT NULL,
  `shopid`   int(10) NOT NULL,
  `cashid`     int(10) NOT NULL,
  `qrid`     int(10) NOT NULL,


  PRIMARY KEY(tableid,shopid,cashid,qrid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
​
CREATE TABLE `orders`(
  `orderid`     int(10) NOT NULL,
  `tableid`   int(10) NOT NULL,
  `shopid`     int(10) NOT NULL,

  PRIMARY KEY(orderid,shopid,tableid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
​
CREATE TABLE `menu`(
  `menuid`     int(10) NOT NULL,
  `shopid`   int(10) NOT NULL,
  `foodid`     int(10) NOT NULL,

  PRIMARY KEY(menuid,shopid,foodid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;