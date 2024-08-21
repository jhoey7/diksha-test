-- MariaDB dump 10.19  Distrib 10.4.21-MariaDB, for osx10.10 (x86_64)
--
-- Host: 127.0.0.1    Database: test-db
-- ------------------------------------------------------
-- Server version	10.4.21-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `orders`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders` (
  `pubid` varchar(100) NOT NULL,
  `order_no` varchar(50) DEFAULT NULL,
  `order_date` timestamp NULL DEFAULT NULL,
  `order_expired` timestamp NULL DEFAULT NULL,
  `user_pubid` varchar(100) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`pubid`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` VALUES ('f1c9d1d4-fd69-4cb2-83c7-a44342df33ad','17EDC5B78759DB20','2024-08-21 07:47:04','2024-08-21 07:48:04','7f8fa80d-0465-47cb-9458-73992cf2e341','CONFIRM','2024-08-21 07:47:04','2024-08-21 07:47:13');

--
-- Table structure for table `orders_details`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders_details` (
  `id` int(11) NOT NULL,
  `order_pubid` varchar(100) DEFAULT NULL,
  `product_pubid` varchar(100) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `orders_details_product_pubid_fk` (`product_pubid`),
  CONSTRAINT `orders_details_product_pubid_fk` FOREIGN KEY (`product_pubid`) REFERENCES `products` (`pubid`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders_details`
--

INSERT INTO `orders_details` VALUES (28,'f1c9d1d4-fd69-4cb2-83c7-a44342df33ad','d6219f47-587b-4719-a585-2627f713012d',8),(29,'f1c9d1d4-fd69-4cb2-83c7-a44342df33ad','e02deae4-7fd4-47b4-8e1c-b07e05d432b2',2);

--
-- Table structure for table `products`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `products` (
  `pubid` varchar(100) NOT NULL,
  `product_code` varchar(50) DEFAULT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `product_description` text DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`pubid`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

INSERT INTO `products` VALUES ('d6219f47-587b-4719-a585-2627f713012d','P-01','Shampoo','Shampoo Test','2024-04-04 21:24:39','2024-04-04 21:24:39'),('e02deae4-7fd4-47b4-8e1c-b07e05d432b2','P-02','Sugar','Sugar Test','2024-04-04 21:24:39','2024-04-04 21:24:39');

--
-- Table structure for table `products_details`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `products_details` (
  `id` int(11) NOT NULL,
  `product_pubid` varchar(100) DEFAULT NULL,
  `wh_id` int(11) DEFAULT NULL,
  `stock` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `product_details_product_pubid_fk` (`product_pubid`),
  KEY `products_details_warehouse_id_fk` (`wh_id`),
  CONSTRAINT `product_details_product_pubid_fk` FOREIGN KEY (`product_pubid`) REFERENCES `products` (`pubid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `products_details_warehouse_id_fk` FOREIGN KEY (`wh_id`) REFERENCES `warehouse` (`id`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products_details`
--

INSERT INTO `products_details` VALUES (1,'d6219f47-587b-4719-a585-2627f713012d',1,0,'2024-04-04 21:26:34','2024-04-04 21:26:34'),(2,'d6219f47-587b-4719-a585-2627f713012d',2,2,'2024-04-04 21:26:34','2024-04-04 21:26:34'),(3,'e02deae4-7fd4-47b4-8e1c-b07e05d432b2',2,3,'2024-04-04 21:26:34','2024-04-04 21:26:34');

--
-- Table structure for table `users`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `pubid` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `mdn` varchar(15) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `username` varchar(25) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `date_of_birth` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `is_unlimited_swipe` tinyint(1) DEFAULT 0,
  `is_verified` tinyint(1) DEFAULT 0,
  `token` text DEFAULT NULL,
  `expire_token_ts` timestamp NULL DEFAULT NULL,
  `counter_password` int(11) DEFAULT 0,
  `counter_password_exp_ts` timestamp NULL DEFAULT NULL,
  `create_ts` timestamp NULL DEFAULT NULL,
  `create_by` varchar(50) DEFAULT NULL,
  `update_ts` timestamp NULL DEFAULT NULL,
  `update_by` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`pubid`),
  KEY `users_mdn_email_index` (`mdn`,`email`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

INSERT INTO `users` VALUES ('7f8fa80d-0465-47cb-9458-73992cf2e341','john.doe@gmail.com','6281787765454','$2a$14$yMN82NWSfclqApMcfzdcDeKpZyAiH4hEshW.bPmCThTkuyFHv5nr2','john','John Doe','Jl. Pangandaran 100','1992-01-01','M',0,0,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQyNDkyNzQsInVzZXJuYW1lIjoiNjI4MTc4Nzc2NTQ1NCJ9.pVezeFFYytwfj-1S52jN8L4wgTztQ9i2ZB-uQR6B3i0','2024-08-21 07:07:54',0,'2024-02-02 18:05:56','2024-02-02 09:11:40','SYSTEM','2024-08-20 07:07:54','SYSTEM'),('ba5325be-546b-44c6-b0a8-e4d3160499b1','john.doe@gmail.com','6281787765456','$2a$14$bIe5HF01e1.34ar6gOvf/u2MN72yekmoECyCfTBCKG0vLXErz5DD2','john','John Doe','Jl. Pangandaran 100','1992-01-01','M',0,0,'',NULL,0,NULL,'2024-04-04 14:12:30','SYSTEM','2024-04-04 14:12:30','');

--
-- Table structure for table `warehouse`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warehouse` (
  `id` int(11) NOT NULL,
  `wh_code` varchar(15) DEFAULT NULL,
  `wh_name` varchar(255) DEFAULT NULL,
  `wh_description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse`
--

INSERT INTO `warehouse` VALUES (1,'WH-01','WH 01','WH 01 Desc','2024-04-04 21:23:39','2024-04-04 21:23:39'),(2,'WH-02','WH 02','WH 02 Desc','2024-04-04 21:23:39','2024-04-04 21:23:39');
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-08-21 22:39:11
