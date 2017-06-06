CREATE TABLE `users` (
     `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
     `name` varchar(30) DEFAULT NULL,
     `email` varchar(50) DEFAULT NULL,
     `password` varchar(60) DEFAULT NULL,
     `remember_token` varchar(100) DEFAULT '',
     `header` varchar(255) DEFAULT '/public/images/header.jpg',
     `sex` tinyint(4) DEFAULT '0',
     `admin` tinyint(4) DEFAULT '0',
     `created_at` int(11) DEFAULT 0,
     `updated_at` int(11) DEFAULT 0,
     PRIMARY KEY (`id`),
     UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


