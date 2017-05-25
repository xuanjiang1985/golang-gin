CREATE TABLE `articles` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	 `item` tinyint(4) DEFAULT '0' COMMENT '0-PHP,1-GO',
     `thanks` int(11) unsigned DEFAULT '0',
     `content` varchar(400) DEFAULT NULL,
     `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;