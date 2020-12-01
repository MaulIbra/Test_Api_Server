-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: db_user
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `m_account`
--

DROP TABLE IF EXISTS `m_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_account` (
  `account_id` varchar(36) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(128) NOT NULL,
  PRIMARY KEY (`account_id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_account`
--

LOCK TABLES `m_account` WRITE;
/*!40000 ALTER TABLE `m_account` DISABLE KEYS */;
INSERT INTO `m_account` VALUES ('1a8d2a2a-63f8-4458-b00a-c543d59544c0','maulanaibrahim@gmail.com','$2a$04$xBU5TUb5xZj.3.ww40m3DuFO5YMWal/URv23gBjh9hlWCOEO2M8L.'),('4eab8034-9cae-4786-a6ca-1b379cf822a5','leaven.cloth@gmail.com','$2a$04$c6sVCrLl4dvBRYRgN.hwEOQ2VC21BsCeNUuj.wvGHXrPpqxPI1JUy'),('612cfb4e-ac40-4b60-b105-c541a90c01ed','maulibrahim19@gmail.com','$2a$04$Ds8vJwrkRxTsxygUYq046.jmVjBSAofa/ADlraaWF0LCrJ6l5Ar4W');
/*!40000 ALTER TABLE `m_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_pekerjaan`
--

DROP TABLE IF EXISTS `m_pekerjaan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_pekerjaan` (
  `id_pekerjaan` varchar(36) NOT NULL,
  `label_pekerjaan` varchar(36) NOT NULL,
  PRIMARY KEY (`id_pekerjaan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_pekerjaan`
--

LOCK TABLES `m_pekerjaan` WRITE;
/*!40000 ALTER TABLE `m_pekerjaan` DISABLE KEYS */;
INSERT INTO `m_pekerjaan` VALUES ('8219c7cc-30bb-11eb-b405-c85b766bafe8','PNS'),('821e309f-30bb-11eb-b405-c85b766bafe8','Wirausaha'),('82225b07-30bb-11eb-b405-c85b766bafe8','Wiraswasta');
/*!40000 ALTER TABLE `m_pekerjaan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_pendidikan`
--

DROP TABLE IF EXISTS `m_pendidikan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_pendidikan` (
  `id_pendidikan` varchar(36) NOT NULL,
  `label_pendidikan` varchar(36) NOT NULL,
  PRIMARY KEY (`id_pendidikan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_pendidikan`
--

LOCK TABLES `m_pendidikan` WRITE;
/*!40000 ALTER TABLE `m_pendidikan` DISABLE KEYS */;
INSERT INTO `m_pendidikan` VALUES ('5ddbb91e-30bb-11eb-b405-c85b766bafe8','SD'),('5de17471-30bb-11eb-b405-c85b766bafe8','SMP'),('5de57c7c-30bb-11eb-b405-c85b766bafe8','SMA'),('5de9a553-30bb-11eb-b405-c85b766bafe8','SD'),('5dedbec5-30bb-11eb-b405-c85b766bafe8','Diploma'),('5df1331d-30bb-11eb-b405-c85b766bafe8','Sarjana'),('5df4cbfb-30bb-11eb-b405-c85b766bafe8','Magister'),('5df8d140-30bb-11eb-b405-c85b766bafe8','Doktor');
/*!40000 ALTER TABLE `m_pendidikan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_user`
--

DROP TABLE IF EXISTS `m_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_user` (
  `id_user` varchar(36) NOT NULL,
  `nik` varchar(16) NOT NULL,
  `username` varchar(45) NOT NULL,
  `tgl_lahir` date NOT NULL,
  `id_pekerjaan` varchar(36) NOT NULL,
  `id_pendidikan` varchar(36) NOT NULL,
  `user_status` int DEFAULT '1',
  `created_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_date` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `nik_UNIQUE` (`nik`),
  KEY `fk_m_user_m_pekerjaan_idx` (`id_pekerjaan`),
  KEY `fk_m_user_m_pendidikan1_idx` (`id_pendidikan`),
  CONSTRAINT `fk_m_user_m_pekerjaan` FOREIGN KEY (`id_pekerjaan`) REFERENCES `m_pekerjaan` (`id_pekerjaan`),
  CONSTRAINT `fk_m_user_m_pendidikan1` FOREIGN KEY (`id_pendidikan`) REFERENCES `m_pendidikan` (`id_pendidikan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_user`
--

LOCK TABLES `m_user` WRITE;
/*!40000 ALTER TABLE `m_user` DISABLE KEYS */;
INSERT INTO `m_user` VALUES ('36d0035f-cd53-4b55-a50c-0ce77d7b8ee4','3209120503980004','MAULANA IBRAHIM','1998-03-05','82225b07-30bb-11eb-b405-c85b766bafe8','5dedbec5-30bb-11eb-b405-c85b766bafe8',1,'2020-12-01 09:36:44','2020-12-01 09:36:44'),('8a95efea-e6d6-47c3-88e7-d57b03f9712e','3209120503980005','UJI COBA','1998-03-05','82225b07-30bb-11eb-b405-c85b766bafe8','5dedbec5-30bb-11eb-b405-c85b766bafe8',1,'2020-12-01 09:41:42','2020-12-01 09:41:42'),('cc45495f-6184-42a1-9a99-8583c8b5f638','320912261197026','PUTRI DIAN INSANI','1997-11-26','82225b07-30bb-11eb-b405-c85b766bafe8','5dedbec5-30bb-11eb-b405-c85b766bafe8',1,'2020-12-01 09:38:52','2020-12-01 09:38:52'),('d8557285-e543-4a14-9fdf-3d7bbbbfc91f','3209120112970004','REZA DWI KURNIAWAN','1997-12-01','8219c7cc-30bb-11eb-b405-c85b766bafe8','5dedbec5-30bb-11eb-b405-c85b766bafe8',1,'2020-12-01 09:37:40','2020-12-01 09:39:14');
/*!40000 ALTER TABLE `m_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-12-01  9:59:46
