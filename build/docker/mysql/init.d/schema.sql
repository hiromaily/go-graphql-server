-- Host: 127.0.0.1    Database: go-graphql
-- ------------------------------------------------------
-- Server version	5.7.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


--
-- DATABASE go-graphql
--
DROP DATABASE IF EXISTS `go-graphql`;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `go-graphql` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `go-graphql`;


--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'user id',
  `name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'name',
  `age` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT'age',
  `country_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT'age',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='user table';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `t_user` WRITE;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
INSERT INTO `t_user` VALUES
  (1,'Dan',24,230,now(),now()),
  (2,'Lee',39,44,now(),now()),
  (3,'Nick',31,229,now(),now());
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;


--
-- Table structure for table `t_company`
--

DROP TABLE IF EXISTS `t_company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_company` (
  `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'company id',
  `name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'company name',
  `country_id` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT'age',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='company table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_company` WRITE;
/*!40000 ALTER TABLE `t_company` DISABLE KEYS */;
INSERT INTO `t_company` VALUES
  (1,'Google',230,now(),now()),
  (2,'Amazon',230,now(),now()),
  (3,'Facebook',230,now(),now()),
  (4,'Apple',230,now(),now()),
  (5,'Netflix',230,now(),now()),
  (6,'AirBnb',230,now(),now()),
  (7,'Twitter',230,now(),now()),
  (8,'Booking.com',155,now(),now()),
  (9,'Transferwise',229,now(),now()),
  (10,'Alibaba',44,now(),now()),
  (11,'Toyota',110,now(),now());
/*!40000 ALTER TABLE `t_company` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_user_work_history`
--

DROP TABLE IF EXISTS `t_user_work_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user_work_history` (
  `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT'id',
  `user_id`     int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'user id',
  `company_id`  int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'company branch id',
  `title`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'title',
  `description` json NOT NULL COMMENT'description',
  `tech_ids`    json NOT NULL COMMENT'tech ids',
  `started_at`  date DEFAULT NULL COMMENT'started date',
  `ended_at`    date DEFAULT NULL COMMENT'ended date',
  `created_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`),
  INDEX user_id (`user_id`),
  INDEX started_at (`started_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='user work history table';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `t_user_work_history` WRITE;
/*!40000 ALTER TABLE `t_user_work_history` DISABLE KEYS */;
INSERT INTO `t_user_work_history` VALUES
(null,1,1,'Backend Developer',
 '["Developed AD server"]',
 '[]',
 '2017-12-1','2021-03-31',now(),now()),
(null,1,2,'Junior Software Engineer',
 '["Developed Game on SmartPhone"]',
 '[]',
 '2010-04-01','2016-11-30',now(),now());
/*!40000 ALTER TABLE `t_user_work_history` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `m_country`
--

CREATE TABLE `m_country` (
  `id`         SMALLINT NOT NULL AUTO_INCREMENT COMMENT'country id',
  `country_code` varchar(2) NOT NULL default'' COMMENT'country code',
  `name`       varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT'country name',
  `created_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
  `updated_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='country table';
/*!40101 SET character_set_client = @saved_cs_client */;
;

LOCK TABLES `m_country` WRITE;
/*!40000 ALTER TABLE `m_country` DISABLE KEYS */;
INSERT INTO `m_country` VALUES 
(null,'AF','Afghanistan',now(),now()),
(null,'AL','Albania',now(),now()),
(null,'DZ','Algeria',now(),now()),
(null,'DS','American Samoa',now(),now()),
(null,'AD','Andorra',now(),now()),
(null,'AO','Angola',now(),now()),
(null,'AI','Anguilla',now(),now()),
(null,'AQ','Antarctica',now(),now()),
(null,'AG','Antigua and Barbuda',now(),now()),
(null,'AR','Argentina',now(),now()),
(null,'AM','Armenia',now(),now()),
(null,'AW','Aruba',now(),now()),
(null,'AU','Australia',now(),now()),
(null,'AT','Austria',now(),now()),
(null,'AZ','Azerbaijan',now(),now()),
(null,'BS','Bahamas',now(),now()),
(null,'BH','Bahrain',now(),now()),
(null,'BD','Bangladesh',now(),now()),
(null,'BB','Barbados',now(),now()),
(null,'BY','Belarus',now(),now()),
(null,'BE','Belgium',now(),now()),
(null,'BZ','Belize',now(),now()),
(null,'BJ','Benin',now(),now()),
(null,'BM','Bermuda',now(),now()),
(null,'BT','Bhutan',now(),now()),
(null,'BO','Bolivia',now(),now()),
(null,'BA','Bosnia and Herzegovina',now(),now()),
(null,'BW','Botswana',now(),now()),
(null,'BV','Bouvet Island',now(),now()),
(null,'BR','Brazil',now(),now()),
(null,'IO','British Indian Ocean Territory',now(),now()),
(null,'BN','Brunei Darussalam',now(),now()),
(null,'BG','Bulgaria',now(),now()),
(null,'BF','Burkina Faso',now(),now()),
(null,'BI','Burundi',now(),now()),
(null,'KH','Cambodia',now(),now()),
(null,'CM','Cameroon',now(),now()),
(null,'CA','Canada',now(),now()),
(null,'CV','Cape Verde',now(),now()),
(null,'KY','Cayman Islands',now(),now()),
(null,'CF','Central African Republic',now(),now()),
(null,'TD','Chad',now(),now()),
(null,'CL','Chile',now(),now()),
(null,'CN','China',now(),now()),
(null,'CX','Christmas Island',now(),now()),
(null,'CC','Cocos (Keeling) Islands',now(),now()),
(null,'CO','Colombia',now(),now()),
(null,'KM','Comoros',now(),now()),
(null,'CG','Congo',now(),now()),
(null,'CK','Cook Islands',now(),now()),
(null,'CR','Costa Rica',now(),now()),
(null,'HR','Croatia (Hrvatska)',now(),now()),
(null,'CU','Cuba',now(),now()),
(null,'CY','Cyprus',now(),now()),
(null,'CZ','Czech Republic',now(),now()),
(null,'DK','Denmark',now(),now()),
(null,'DJ','Djibouti',now(),now()),
(null,'DM','Dominica',now(),now()),
(null,'DO','Dominican Republic',now(),now()),
(null,'TP','East Timor',now(),now()),
(null,'EC','Ecuador',now(),now()),
(null,'EG','Egypt',now(),now()),
(null,'SV','El Salvador',now(),now()),
(null,'GQ','Equatorial Guinea',now(),now()),
(null,'ER','Eritrea',now(),now()),
(null,'EE','Estonia',now(),now()),
(null,'ET','Ethiopia',now(),now()),
(null,'FK','Falkland Islands (Malvinas)',now(),now()),
(null,'FO','Faroe Islands',now(),now()),
(null,'FJ','Fiji',now(),now()),
(null,'FI','Finland',now(),now()),
(null,'FR','France',now(),now()),
(null,'FX','France, Metropolitan',now(),now()),
(null,'GF','French Guiana',now(),now()),
(null,'PF','French Polynesia',now(),now()),
(null,'TF','French Southern Territories',now(),now()),
(null,'GA','Gabon',now(),now()),
(null,'GM','Gambia',now(),now()),
(null,'GE','Georgia',now(),now()),
(null,'DE','Germany',now(),now()),
(null,'GH','Ghana',now(),now()),
(null,'GI','Gibraltar',now(),now()),
(null,'GK','Guernsey',now(),now()),
(null,'GR','Greece',now(),now()),
(null,'GL','Greenland',now(),now()),
(null,'GD','Grenada',now(),now()),
(null,'GP','Guadeloupe',now(),now()),
(null,'GU','Guam',now(),now()),
(null,'GT','Guatemala',now(),now()),
(null,'GN','Guinea',now(),now()),
(null,'GW','Guinea-Bissau',now(),now()),
(null,'GY','Guyana',now(),now()),
(null,'HT','Haiti',now(),now()),
(null,'HM','Heard and Mc Donald Islands',now(),now()),
(null,'HN','Honduras',now(),now()),
(null,'HK','Hong Kong',now(),now()),
(null,'HU','Hungary',now(),now()),
(null,'IS','Iceland',now(),now()),
(null,'IN','India',now(),now()),
(null,'IM','Isle of Man',now(),now()),
(null,'ID','Indonesia',now(),now()),
(null,'IR','Iran (Islamic Republic of)',now(),now()),
(null,'IQ','Iraq',now(),now()),
(null,'IE','Ireland',now(),now()),
(null,'IL','Israel',now(),now()),
(null,'IT','Italy',now(),now()),
(null,'CI','Ivory Coast',now(),now()),
(null,'JE','Jersey',now(),now()),
(null,'JM','Jamaica',now(),now()),
(null,'JP','Japan',now(),now()),
(null,'JO','Jordan',now(),now()),
(null,'KZ','Kazakhstan',now(),now()),
(null,'KE','Kenya',now(),now()),
(null,'KI','Kiribati',now(),now()),
(null,'KP','Korea, Democratic People''s Republic of',now(),now()),
(null,'KR','Korea, Republic of',now(),now()),
(null,'XK','Kosovo',now(),now()),
(null,'KW','Kuwait',now(),now()),
(null,'KG','Kyrgyzstan',now(),now()),
(null,'LA','Lao People''s Democratic Republic',now(),now()),
(null,'LV','Latvia',now(),now()),
(null,'LB','Lebanon',now(),now()),
(null,'LS','Lesotho',now(),now()),
(null,'LR','Liberia',now(),now()),
(null,'LY','Libyan Arab Jamahiriya',now(),now()),
(null,'LI','Liechtenstein',now(),now()),
(null,'LT','Lithuania',now(),now()),
(null,'LU','Luxembourg',now(),now()),
(null,'MO','Macau',now(),now()),
(null,'MK','Macedonia',now(),now()),
(null,'MG','Madagascar',now(),now()),
(null,'MW','Malawi',now(),now()),
(null,'MY','Malaysia',now(),now()),
(null,'MV','Maldives',now(),now()),
(null,'ML','Mali',now(),now()),
(null,'MT','Malta',now(),now()),
(null,'MH','Marshall Islands',now(),now()),
(null,'MQ','Martinique',now(),now()),
(null,'MR','Mauritania',now(),now()),
(null,'MU','Mauritius',now(),now()),
(null,'TY','Mayotte',now(),now()),
(null,'MX','Mexico',now(),now()),
(null,'FM','Micronesia, Federated States of',now(),now()),
(null,'MD','Moldova, Republic of',now(),now()),
(null,'MC','Monaco',now(),now()),
(null,'MN','Mongolia',now(),now()),
(null,'ME','Montenegro',now(),now()),
(null,'MS','Montserrat',now(),now()),
(null,'MA','Morocco',now(),now()),
(null,'MZ','Mozambique',now(),now()),
(null,'MM','Myanmar',now(),now()),
(null,'NA','Namibia',now(),now()),
(null,'NR','Nauru',now(),now()),
(null,'NP','Nepal',now(),now()),
(null,'NL','Netherlands',now(),now()),
(null,'AN','Netherlands Antilles',now(),now()),
(null,'NC','New Caledonia',now(),now()),
(null,'NZ','New Zealand',now(),now()),
(null,'NI','Nicaragua',now(),now()),
(null,'NE','Niger',now(),now()),
(null,'NG','Nigeria',now(),now()),
(null,'NU','Niue',now(),now()),
(null,'NF','Norfolk Island',now(),now()),
(null,'MP','Northern Mariana Islands',now(),now()),
(null,'NO','Norway',now(),now()),
(null,'OM','Oman',now(),now()),
(null,'PK','Pakistan',now(),now()),
(null,'PW','Palau',now(),now()),
(null,'PS','Palestine',now(),now()),
(null,'PA','Panama',now(),now()),
(null,'PG','Papua New Guinea',now(),now()),
(null,'PY','Paraguay',now(),now()),
(null,'PE','Peru',now(),now()),
(null,'PH','Philippines',now(),now()),
(null,'PN','Pitcairn',now(),now()),
(null,'PL','Poland',now(),now()),
(null,'PT','Portugal',now(),now()),
(null,'PR','Puerto Rico',now(),now()),
(null,'QA','Qatar',now(),now()),
(null,'RE','Reunion',now(),now()),
(null,'RO','Romania',now(),now()),
(null,'RU','Russian Federation',now(),now()),
(null,'RW','Rwanda',now(),now()),
(null,'KN','Saint Kitts and Nevis',now(),now()),
(null,'LC','Saint Lucia',now(),now()),
(null,'VC','Saint Vincent and the Grenadines',now(),now()),
(null,'WS','Samoa',now(),now()),
(null,'SM','San Marino',now(),now()),
(null,'ST','Sao Tome and Principe',now(),now()),
(null,'SA','Saudi Arabia',now(),now()),
(null,'SN','Senegal',now(),now()),
(null,'RS','Serbia',now(),now()),
(null,'SC','Seychelles',now(),now()),
(null,'SL','Sierra Leone',now(),now()),
(null,'SG','Singapore',now(),now()),
(null,'SK','Slovakia',now(),now()),
(null,'SI','Slovenia',now(),now()),
(null,'SB','Solomon Islands',now(),now()),
(null,'SO','Somalia',now(),now()),
(null,'ZA','South Africa',now(),now()),
(null,'GS','South Georgia South Sandwich Islands',now(),now()),
(null,'ES','Spain',now(),now()),
(null,'LK','Sri Lanka',now(),now()),
(null,'SH','St. Helena',now(),now()),
(null,'PM','St. Pierre and Miquelon',now(),now()),
(null,'SD','Sudan',now(),now()),
(null,'SR','Suriname',now(),now()),
(null,'SJ','Svalbard and Jan Mayen Islands',now(),now()),
(null,'SZ','Swaziland',now(),now()),
(null,'SE','Sweden',now(),now()),
(null,'CH','Switzerland',now(),now()),
(null,'SY','Syrian Arab Republic',now(),now()),
(null,'TW','Taiwan',now(),now()),
(null,'TJ','Tajikistan',now(),now()),
(null,'TZ','Tanzania',now(),now()),
(null,'TH','Thailand',now(),now()),
(null,'TG','Togo',now(),now()),
(null,'TK','Tokelau',now(),now()),
(null,'TO','Tonga',now(),now()),
(null,'TT','Trinidad and Tobago',now(),now()),
(null,'TN','Tunisia',now(),now()),
(null,'TR','Turkey',now(),now()),
(null,'TM','Turkmenistan',now(),now()),
(null,'TC','Turks and Caicos Islands',now(),now()),
(null,'TV','Tuvalu',now(),now()),
(null,'UG','Uganda',now(),now()),
(null,'UA','Ukraine',now(),now()),
(null,'AE','United Arab Emirates',now(),now()),
(null,'GB','United Kingdom',now(),now()),
(null,'US','United States',now(),now()),
(null,'UM','United States minor outlying islands',now(),now()),
(null,'UY','Uruguay',now(),now()),
(null,'UZ','Uzbekistan',now(),now()),
(null,'VU','Vanuatu',now(),now()),
(null,'VA','Vatican City State',now(),now()),
(null,'VE','Venezuela',now(),now()),
(null,'VN','Vietnam',now(),now()),
(null,'VG','Virgin Islands (British)',now(),now()),
(null,'VI','Virgin Islands (U.S.)',now(),now()),
(null,'WF','Wallis and Futuna Islands',now(),now()),
(null,'EH','Western Sahara',now(),now()),
(null,'YE','Yemen',now(),now()),
(null,'ZR','Zaire',now(),now()),
(null,'ZM','Zambia',now(),now()),
(null,'ZW','Zimbabwe',now(),now());
/*!40000 ALTER TABLE `m_country` ENABLE KEYS */;
UNLOCK TABLES;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
