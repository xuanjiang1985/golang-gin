CREATE TABLE `articles` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	 `user_id` int(11) unsigned DEFAULT '0' COMMENT '0-no',
     `thanks` int(11) unsigned DEFAULT '0',
     `comments` int(11) unsigned DEFAULT '0',
     `content` varchar(400) DEFAULT NULL,
     `created_at` int(11) DEFAULT 0,
     `updated_at` int(11) DEFAULT 0,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;