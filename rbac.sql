/*
MySQL - 5.6.17 : Database - rbac
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`rbac` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `rbac`;

/*Table structure for table `rbac_access` */

DROP TABLE IF EXISTS `rbac_access`;

CREATE TABLE `rbac_access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` smallint(6) unsigned NOT NULL,
  `node_id` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) NOT NULL,
  `module` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `groupId` (`role_id`),
  KEY `nodeId` (`node_id`)
) ENGINE=MyISAM AUTO_INCREMENT=79 DEFAULT CHARSET=utf8;

/*Data for the table `rbac_access` */

insert  into `rbac_access`(`id`,`role_id`,`node_id`,`level`,`module`) values (77,1,10,3,'操作'),(76,1,4,2,'模块'),(75,1,12,3,'操作'),(74,1,9,3,'操作'),(73,1,3,2,'模块'),(72,1,13,3,'操作'),(71,1,8,3,'操作'),(70,1,6,3,'操作'),(69,1,5,3,'操作'),(68,1,2,2,'模块'),(67,1,1,1,'项目'),(64,15,9,3,'操作'),(63,15,3,2,'模块'),(62,15,8,3,'操作'),(33,2,12,3,'操作'),(32,2,9,3,'操作'),(31,2,3,2,'模块'),(30,2,1,1,'项目'),(61,15,6,3,'操作'),(60,15,2,2,'模块'),(59,15,1,1,'项目'),(65,15,12,3,'操作'),(66,15,7,2,'模块'),(78,1,7,2,'模块');

/*Table structure for table `rbac_node` */

DROP TABLE IF EXISTS `rbac_node`;

CREATE TABLE `rbac_node` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `sort` smallint(6) unsigned DEFAULT NULL,
  `pid` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) unsigned NOT NULL,
  `is_show` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `level` (`level`),
  KEY `pid` (`pid`),
  KEY `status` (`status`),
  KEY `name` (`name`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

/*Data for the table `rbac_node` */

insert  into `rbac_node`(`id`,`name`,`title`,`status`,`remark`,`sort`,`pid`,`level`,`is_show`) values (1,'RBAC','后台管理',1,'',1,0,1,1),(2,'User','用户管理',1,'',1,1,2,1),(3,'Hack','漏洞管理',1,'',2,1,2,1),(4,'Week','周报管理',1,'',3,1,2,1),(5,'RoleList','角色管理',1,'',1,2,3,1),(6,'NodeList','权限管理',1,'',2,2,3,1),(7,'Auto','文章管理',1,'',4,1,2,1),(8,'UserList','用户管理',1,'',3,2,3,1),(9,'hack-list','漏洞列表',1,'',1,3,3,1),(10,'week-list','周报列表',1,'',1,4,3,1),(12,'hack-json','漏洞数据',1,'',2,3,3,1),(13,'RoleListJson','角色列表',1,'',4,2,3,0);

/*Table structure for table `rbac_role` */

DROP TABLE IF EXISTS `rbac_role`;

CREATE TABLE `rbac_role` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `pid` smallint(6) DEFAULT NULL,
  `status` tinyint(1) unsigned DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `status` (`status`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

/*Data for the table `rbac_role` */

insert  into `rbac_role`(`id`,`name`,`pid`,`status`,`remark`) values (1,'超级管理员',0,1,''),(2,'普通会员',1,1,''),(3,'普通用户',1,1,''),(15,'漏洞审核部门',1,1,'');

/*Table structure for table `rbac_role_user` */

DROP TABLE IF EXISTS `rbac_role_user`;

CREATE TABLE `rbac_role_user` (
  `role_id` mediumint(9) unsigned NOT NULL,
  `user_id` int(5) NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `group_id` (`role_id`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `rbac_role_user` */

insert  into `rbac_role_user`(`role_id`,`user_id`) values (1,1),(3,4),(2,3),(15,5);

/*Table structure for table `rbac_user` */

DROP TABLE IF EXISTS `rbac_user`;

CREATE TABLE `rbac_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(60) NOT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `loginip` varchar(17) DEFAULT NULL,
  `logintime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

/*Data for the table `rbac_user` */

insert  into `rbac_user`(`id`,`username`,`password`,`status`,`loginip`,`logintime`) values (1,'admin','admin',1,'','0000-00-00 00:00:00'),(3,'a','a',1,'','0000-00-00 00:00:00'),(4,'t','t',1,'','0000-00-00 00:00:00'),(5,'b','b',1,'','0000-00-00 00:00:00');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
