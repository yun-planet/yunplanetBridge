
CREATE DATABASE `yunplanet_watchjob` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

CREATE TABLE `address_checkjob` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `user_address_id` int(11) DEFAULT NULL,
  `limitcount` int(255) DEFAULT NULL,
  `jobtype` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;