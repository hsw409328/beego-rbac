# ************************************************************
# Sequel Pro SQL dump
# Version 5438
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 8.0.23)
# Database: rbac
# Generation Time: 2022-08-28 06:37:39 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table rbac_access
# ------------------------------------------------------------

DROP TABLE IF EXISTS `rbac_access`;

CREATE TABLE `rbac_access` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` smallint unsigned NOT NULL,
  `node_id` smallint unsigned NOT NULL,
  `level` tinyint(1) NOT NULL,
  `module` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `groupId` (`role_id`),
  KEY `nodeId` (`node_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

LOCK TABLES `rbac_access` WRITE;
/*!40000 ALTER TABLE `rbac_access` DISABLE KEYS */;

INSERT INTO `rbac_access` (`id`, `role_id`, `node_id`, `level`, `module`)
VALUES
	(77,1,10,3,'操作'),
	(76,1,4,2,'模块'),
	(75,1,12,3,'操作'),
	(74,1,9,3,'操作'),
	(73,1,3,2,'模块'),
	(72,1,13,3,'操作'),
	(71,1,8,3,'操作'),
	(70,1,6,3,'操作'),
	(69,1,5,3,'操作'),
	(68,1,2,2,'模块'),
	(67,1,1,1,'项目'),
	(64,15,9,3,'操作'),
	(63,15,3,2,'模块'),
	(62,15,8,3,'操作'),
	(33,2,12,3,'操作'),
	(32,2,9,3,'操作'),
	(31,2,3,2,'模块'),
	(30,2,1,1,'项目'),
	(61,15,6,3,'操作'),
	(60,15,2,2,'模块'),
	(59,15,1,1,'项目'),
	(65,15,12,3,'操作'),
	(66,15,7,2,'模块'),
	(78,1,7,2,'模块'),
	(96,16,7,2,'模块'),
	(95,16,13,3,'操作'),
	(94,16,5,3,'操作'),
	(93,16,2,2,'模块'),
	(92,16,1,1,'项目');

/*!40000 ALTER TABLE `rbac_access` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table rbac_node
# ------------------------------------------------------------

DROP TABLE IF EXISTS `rbac_node`;

CREATE TABLE `rbac_node` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `sort` smallint unsigned DEFAULT NULL,
  `pid` smallint unsigned NOT NULL,
  `level` tinyint unsigned NOT NULL,
  `is_show` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `level` (`level`),
  KEY `pid` (`pid`),
  KEY `status` (`status`),
  KEY `name` (`name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

LOCK TABLES `rbac_node` WRITE;
/*!40000 ALTER TABLE `rbac_node` DISABLE KEYS */;

INSERT INTO `rbac_node` (`id`, `name`, `title`, `status`, `remark`, `sort`, `pid`, `level`, `is_show`)
VALUES
	(1,'RBAC','后台管理',1,'',1,0,1,1),
	(2,'/User','用户管理',1,'',1,1,2,1),
	(3,'/Hack','漏洞管理',1,'',2,1,2,1),
	(4,'/Week','周报管理',1,'',3,1,2,1),
	(5,'/RoleList','角色管理',1,'',1,2,3,1),
	(6,'/NodeList','权限管理',1,'',2,2,3,1),
	(7,'/Auto','文章管理',1,'',4,1,2,0),
	(8,'/UserList','用户管理',1,'',3,2,3,1),
	(9,'/hack-list','漏洞列表',1,'',1,3,3,1),
	(10,'/week-list','周报列表',1,'',1,4,3,1),
	(12,'/hack-json','漏洞数据',1,'',2,3,3,1),
	(13,'/RoleListJson','角色列表API',1,'',4,2,3,0),
	(14,'/RoleAdd','添加角色API',1,'',5,2,3,0),
	(15,'/RoleDelete','删除角色API',1,'',6,2,3,0),
	(16,'/NodeListJson','权限列表API',1,'',7,2,3,0),
	(17,'/NodeAdd','权限列表API',1,'',8,2,3,0),
	(18,'/NodeDelete','权限列表API',1,'',9,2,3,0),
	(19,'/UserListJson','权限列表API',1,'',10,2,3,0),
	(20,'/UserAdd','权限列表API',1,'',11,2,3,0),
	(21,'/UserDelete','权限列表API',1,'',12,2,3,0),
	(22,'/AccessListJson','配置权限列表API',1,'',13,2,3,0),
	(23,'/AccessAdd','配置权限API',1,'',14,2,3,0);

/*!40000 ALTER TABLE `rbac_node` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table rbac_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `rbac_role`;

CREATE TABLE `rbac_role` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `pid` smallint DEFAULT NULL,
  `status` tinyint unsigned DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `status` (`status`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

LOCK TABLES `rbac_role` WRITE;
/*!40000 ALTER TABLE `rbac_role` DISABLE KEYS */;

INSERT INTO `rbac_role` (`id`, `name`, `pid`, `status`, `remark`)
VALUES
	(1,'超级管理员',0,1,''),
	(2,'普通会员',1,1,''),
	(3,'普通用户',1,1,''),
	(15,'漏洞审核部门',1,1,''),
	(16,'测试权限',0,1,'');

/*!40000 ALTER TABLE `rbac_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table rbac_role_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `rbac_role_user`;

CREATE TABLE `rbac_role_user` (
  `role_id` mediumint unsigned NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `group_id` (`role_id`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

LOCK TABLES `rbac_role_user` WRITE;
/*!40000 ALTER TABLE `rbac_role_user` DISABLE KEYS */;

INSERT INTO `rbac_role_user` (`role_id`, `user_id`)
VALUES
	(1,1),
	(3,4),
	(2,3),
	(15,5),
	(16,6),
	(1,7);

/*!40000 ALTER TABLE `rbac_role_user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table rbac_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `rbac_user`;

CREATE TABLE `rbac_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(60) NOT NULL,
  `status` tinyint DEFAULT NULL,
  `loginip` varchar(17) DEFAULT NULL,
  `logintime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `rbac_user` WRITE;
/*!40000 ALTER TABLE `rbac_user` DISABLE KEYS */;

INSERT INTO `rbac_user` (`id`, `username`, `password`, `status`, `loginip`, `logintime`)
VALUES
	(1,'admin','admin',1,'','0000-00-00 00:00:00'),
	(3,'a','a',1,'','0000-00-00 00:00:00'),
	(4,'t','t',1,'','0000-00-00 00:00:00'),
	(5,'b','b',1,'','0000-00-00 00:00:00'),
	(6,'c','c',1,'','2022-08-28 13:28:25'),
	(7,'e','e',1,'','2022-08-28 14:36:00');

/*!40000 ALTER TABLE `rbac_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
