CREATE USER IF NOT EXISTS 'dev'@'%' IDENTIFIED BY 'dev';
CREATE DATABASE IF NOT EXISTS ur_v2;
GRANT ALL PRIVILEGES ON ur_v2.* TO 'dev'@'%';

USE ur_v2;

CREATE TABLE `prefs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `region` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_crawl` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `prefs_code_unique` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `houses` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `pref_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rooms_got_at` datetime NOT NULL DEFAULT '1000-01-01 00:00:00',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `houses_code_unique` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `rooms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `house_code` varchar(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `room_code` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('ready','closed') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'ready',
  `got_at` datetime NOT NULL DEFAULT '1000-01-01 00:00:00',
  `data` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `rooms_house_code_room_code_unique` (`house_code`,`room_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET NAMES 'utf8mb4';

INSERT INTO `prefs` (`code`, `region`, `name`, `is_crawl`, `created_at`, `updated_at`) VALUES
('hokkaido', 'hokkaitohoku', '北海道', 0, NOW(), NOW()),
('miyagi', 'hokkaitohoku', '宮城県', 0, NOW(), NOW()),
('tokyo', 'kanto', '東京都', 1, NOW(), NOW()),
('kanagawa', 'kanto', '神奈川県', 1, NOW(), NOW()),
('chiba', 'kanto', '千葉県', 1, NOW(), NOW()),
('saitama', 'kanto', '埼玉県', 1, NOW(), NOW()),
('ibaraki', 'kanto', '茨城県', 0, NOW(), NOW()),
('aichi', 'tokai', '愛知県', 0, NOW(), NOW()),
('mie', 'tokai', '三重県', 0, NOW(), NOW()),
('gifu', 'tokai', '岐阜県', 0, NOW(), NOW()),
('shizuoka', 'tokai', '静岡県', 0, NOW(), NOW()),
('osaka', 'kansai', '大阪府', 1, NOW(), NOW()),
('hyogo', 'kansai', '兵庫県', 0, NOW(), NOW()),
('kyoto', 'kansai', '京都府', 0, NOW(), NOW()),
('shiga', 'kansai', '滋賀県', 0, NOW(), NOW()),
('nara', 'kansai', '奈良県', 0, NOW(), NOW()),
('wakayama', 'kansai', '和歌山県', 0, NOW(), NOW()),
('okayama', 'chugoku', '岡山県', 0, NOW(), NOW()),
('hiroshima', 'chugoku', '広島県', 0, NOW(), NOW()),
('yamaguchi', 'chugoku', '山口県', 0, NOW(), NOW()),
('fukuoka', 'kyushu', '福岡県', 1, NOW(), NOW());