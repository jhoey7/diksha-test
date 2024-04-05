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
  PRIMARY KEY (`pubid`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` VALUES ('31969555-a2c1-40bb-b93b-40ec73e9a53f','17C33B5528704DE0','2024-04-04 17:38:40',NULL,'7f8fa80d-0465-47cb-9458-73992cf2e341','PENDING','2024-04-04 17:38:40'),('bd3efcce-3d15-44cb-97f9-32bd53cbb13b','17C33CB24A1A3C18','2024-04-04 18:03:39',NULL,'7f8fa80d-0465-47cb-9458-73992cf2e341','PENDING','2024-04-04 18:03:39');

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

INSERT INTO `orders_details` VALUES (1,'31969555-a2c1-40bb-b93b-40ec73e9a53f','d6219f47-587b-4719-a585-2627f713012d',1),(2,'bd3efcce-3d15-44cb-97f9-32bd53cbb13b','d6219f47-587b-4719-a585-2627f713012d',1);

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
  `total_stock` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`pubid`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

INSERT INTO `products` VALUES ('d6219f47-587b-4719-a585-2627f713012d','P-01','Shampoo','Shampoo Test',9,'2024-04-04 21:24:39','2024-04-04 21:24:39'),('e02deae4-7fd4-47b4-8e1c-b07e05d432b2','P-02','Sugar','Sugar Test',5,'2024-04-04 21:24:39','2024-04-04 21:24:39');

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

INSERT INTO `products_details` VALUES (1,'d6219f47-587b-4719-a585-2627f713012d',1,6,'2024-04-04 21:26:34','2024-04-04 21:26:34'),(2,'d6219f47-587b-4719-a585-2627f713012d',2,4,'2024-04-04 21:26:34','2024-04-04 21:26:34'),(3,'e02deae4-7fd4-47b4-8e1c-b07e05d432b2',2,5,'2024-04-04 21:26:34','2024-04-04 21:26:34');

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

-- Dump completed on 2024-04-05  8:11:32
