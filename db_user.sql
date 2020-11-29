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
INSERT INTO `m_user` VALUES ('5215fb19-5c71-449e-a707-f772239c151b','3209120503980000','Maulana Ibrahim','1998-03-05','821e309f-30bb-11eb-b405-c85b766bafe8','5dedbec5-30bb-11eb-b405-c85b766bafe8',1,'2020-11-29 10:02:52','2020-11-30 06:38:34'),('744635c3-a7fb-4028-a80a-7a85021f0f21','320192020398440','Putri Dian Insani','1998-12-03','82225b07-30bb-11eb-b405-c85b766bafe8','5dedbec5-30bb-11eb-b405-c85b766bafe8',1,'2020-11-29 10:04:19','2020-11-29 10:04:19'),('745cebc1-c128-4a3f-8f77-d96fdf7607f9','3209120503980002','Reza Dwi Kurniawan','1997-12-01','82225b07-30bb-11eb-b405-c85b766bafe8','5dedbec5-30bb-11eb-b405-c85b766bafe8',1,'2020-11-29 10:03:22','2020-11-29 10:03:22');
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

-- Dump completed on 2020-11-30  6:40:49
