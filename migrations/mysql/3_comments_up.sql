CREATE TABLE `comments` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) unsigned DEFAULT '0',
  `user_id` int(11) unsigned DEFAULT '0',
  `comment` varchar(255) DEFAULT NULL,
  `created_at` int(11) DEFAULT 0,
  `updated_at` int(11) DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `comments_article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;