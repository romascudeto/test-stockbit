-- `go`.logs definition

CREATE TABLE `logs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `activity` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT '""',
  `created_at` datetime DEFAULT NULL,
  `request` text COLLATE utf8_unicode_ci,
  `response` text COLLATE utf8_unicode_ci,
  `url` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL,
  `response_status` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;