CREATE DATABASE `yun_node` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

CREATE TABLE `transact` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `addressTo` varchar(56) DEFAULT NULL,
  `blockNumber` varchar(56) DEFAULT NULL,
  `txtHash` varchar(72) DEFAULT NULL,
  `coinType` varchar(11) DEFAULT NULL COMMENT 'eth,omg,yun',
  `coinValue` varchar(72) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `transType` int(11) DEFAULT NULL COMMENT '1.send eth 2.erc20 3.create contract',
  `coinNumber` decimal(19,18) DEFAULT '0.000000000000000000',
  `addressFrom` varchar(56) DEFAULT NULL,
  `handerStatus` int(11) DEFAULT '0',
  `transtatus` int(11) DEFAULT '0' COMMENT 'trade status',
  `coinValue2` varchar(36) DEFAULT NULL COMMENT 'text',
  `nonce` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
