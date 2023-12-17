-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.35-0ubuntu0.22.04.1 - (Ubuntu)
-- Server OS:                    Linux
-- HeidiSQL Version:             12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for dbo
DROP DATABASE IF EXISTS `dbo`;
CREATE DATABASE IF NOT EXISTS `dbo` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `dbo`;

-- Dumping structure for table dbo.customers
DROP TABLE IF EXISTS `customers`;
CREATE TABLE IF NOT EXISTS `customers` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `email` longtext,
  `address` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table dbo.customers: ~2 rows (approximately)
INSERT INTO `customers` (`id`, `name`, `email`, `address`) VALUES
	(1, 'ilham', 'ilham@getnada.com', NULL),
	(4, 'rahman', 'rahman@getnada.com', 'Malang');

-- Dumping structure for table dbo.orders
DROP TABLE IF EXISTS `orders`;
CREATE TABLE IF NOT EXISTS `orders` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `transaction_number` varchar(100) DEFAULT NULL,
  `total_price` float DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table dbo.orders: ~2 rows (approximately)
INSERT INTO `orders` (`id`, `transaction_number`, `total_price`, `created_at`, `created_by`) VALUES
	(1, '1234567', 20000, '2023-12-17 06:58:58', 1),
	(8, 'qwewert1234', 30000, '2023-12-17 08:15:11', 1);

-- Dumping structure for table dbo.order_details
DROP TABLE IF EXISTS `order_details`;
CREATE TABLE IF NOT EXISTS `order_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `orders_id` bigint NOT NULL DEFAULT '0',
  `item_name` varchar(50) DEFAULT NULL,
  `quantity` int NOT NULL DEFAULT '0',
  `price` float NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table dbo.order_details: ~5 rows (approximately)
INSERT INTO `order_details` (`id`, `orders_id`, `item_name`, `quantity`, `price`) VALUES
	(1, 1, 'tas', 1, 10000),
	(3, 1, 'pulpen', 2, 5000),
	(5, 8, 'pensil', 2, 2000),
	(6, 8, 'penghapus', 5, 3000),
	(7, 8, 'penggaris', 2, 11000);

-- Dumping structure for table dbo.users
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table dbo.users: ~0 rows (approximately)
INSERT INTO `users` (`id`, `email`, `password`) VALUES
	(1, 'aaa@a.com', '$2a$10$banDwk07kB3GkHP3bkMmf.H2cCHh81h8AicEJc3bE7LwT.36t1zHG');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
