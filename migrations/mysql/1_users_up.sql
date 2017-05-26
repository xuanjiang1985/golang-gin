CREATE TABLE `users` (
     `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
     `name` varchar(30) DEFAULT NULL,
     `email` varchar(50) DEFAULT NULL,
     `password` varchar(60) DEFAULT NULL,
     `remember_token` varchar(100) DEFAULT NULL,
     `sex` tinyint(4) DEFAULT '0',
     `admin` tinyint(4) DEFAULT '0',
     `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


