/*
SQLyog Ultimate v12.5.0 (64 bit)
MySQL - 5.7.21 : Database - book_store
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*Table structure for table `author` */

DROP TABLE IF EXISTS `author`;

CREATE TABLE `author` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `years_active` date NOT NULL,
  `slug` varchar(30) NOT NULL,
  `dob` date NOT NULL,
  `about` varchar(300) DEFAULT NULL,
  `language` varchar(45) DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name_UNIQUE` (`slug`),
  KEY `ind_author_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

/*Data for the table `author` */

insert  into `author`(`id`,`name`,`years_active`,`slug`,`dob`,`about`,`language`,`status`,`created_at`,`updated_at`) values 
(2,'George Orwell','1940-11-12','george-orwell','1899-02-27','True, Orwell (whose real name was Eric Arthur Blair) isn\'t everyone\'s taste, especially those who do not share his views on totalitarianism.But `Animal Farm` and `1984` are exemplary novels that truly get the reader to think more about politics, society and culture..','English',1,'2020-03-01 11:11:32','2020-03-01 11:11:32'),
(3,'J.K. Rowling','1985-09-12','j.k.-rowling','1940-02-27','Like her or not, Ms. Rowling has a style of writing that has launched her into the annals of literary history. Her Harry Potter books have won awards not only for their imagination but also for their strong prose.','English',1,'2020-03-01 11:12:13','2020-03-01 11:12:13'),
(4,'James Joyce','1990-01-12','james-joyce','1940-02-27','Did you read `Ulysses` in school or while at university?Plenty of students did, but most would do well to revisit Joyce\'s most renowned work. Time will not have changed the words, but it makes all the difference in the interpretation','English',1,'2020-03-01 11:12:59','2020-03-01 11:12:59'),
(5,'Satya Vyas','2014-01-12','satya-vyas','1970-06-17','Satya Vyas is professionals-turned-amateur writer of modern hindi also known as “Nai wali Hindi”. He is law graduate from the prestigious law school BHU, and a logistics professional.','Hindi',1,'2020-03-01 11:13:49','2020-03-01 11:13:49'),
(6,'Ramdhari Singh Dinkar','1974-03-26','ramdhari-singh-dinkar','1908-09-23','Ramdhari Singh, known by his nom de plume Dinkar, was an Indian Hindi poet, essayist, patriot and academic, who is considered as one of the most important modern Hindi poets.','Hindi',1,'2020-03-01 11:24:57','2020-03-01 11:24:57'),
(7,'William Shakespeare','1860-03-26','william-shakespeare','1808-12-12','Chances are strong that you\'ve read or seen at least one of his plays, but if it\'s been a while since you perused a copy of `The Tragedy of Hamlet` or `Macbeth`, it\'s time to reacquaint yourself with this master of language and storytelling.','English',1,'2020-03-01 13:57:14','2020-03-01 08:30:13');

/*Table structure for table `book_author` */

DROP TABLE IF EXISTS `book_author`;

CREATE TABLE `book_author` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_id` int(11) NOT NULL,
  `author_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FOR_BookAuthor_BookID_idx` (`book_id`),
  KEY `FOR_BookAuthor_AuthorID_idx` (`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;

/*Data for the table `book_author` */

insert  into `book_author`(`id`,`book_id`,`author_id`,`created_at`) values 
(11,8,2,'2020-03-03 12:21:22'),
(12,8,7,'2020-03-03 12:21:22'),
(13,10,3,'2020-03-03 12:29:04');

/*Table structure for table `book_genre` */

DROP TABLE IF EXISTS `book_genre`;

CREATE TABLE `book_genre` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_id` int(11) NOT NULL,
  `genre_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FOR_BookGenre_BookID_idx` (`book_id`),
  KEY `FOR_BookGenre_GenreID_idx` (`genre_id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=latin1;

/*Data for the table `book_genre` */

insert  into `book_genre`(`id`,`book_id`,`genre_id`,`created_at`) values 
(49,8,5,'2020-03-03 12:21:22'),
(50,8,13,'2020-03-03 12:21:22'),
(51,8,7,'2020-03-03 12:21:22'),
(52,8,6,'2020-03-03 12:21:22'),
(53,10,6,'2020-03-03 12:29:04'),
(54,10,7,'2020-03-03 12:29:04'),
(55,10,8,'2020-03-03 12:29:04'),
(56,10,9,'2020-03-03 12:29:04'),
(57,10,12,'2020-03-03 12:29:04');

/*Table structure for table `books` */

DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `isbn` varchar(45) NOT NULL,
  `price` float NOT NULL,
  `langauge` varchar(45) NOT NULL,
  `quantity` int(11) NOT NULL,
  `old_price` float NOT NULL,
  `book_type` enum('soft','hard') NOT NULL DEFAULT 'hard',
  `publisher_id` int(11) DEFAULT NULL,
  `image` varchar(200) DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `number_pages` int(11) NOT NULL,
  `published_at` date NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `isbn_UNIQUE` (`isbn`),
  KEY `book_name_IND` (`name`),
  KEY `Book_Publisher_id_idx` (`publisher_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

/*Data for the table `books` */

insert  into `books`(`id`,`name`,`isbn`,`price`,`langauge`,`quantity`,`old_price`,`book_type`,`publisher_id`,`image`,`status`,`number_pages`,`published_at`,`created_at`,`updated_at`) values 
(8,'Romeo and Juliet','1586638459',2123.12,'English',100,780,'hard',2,'https://images-na.ssl-images-amazon.com/images/I/414ox%2BURkdL.jpg',1,120,'2010-08-08','2020-03-02 11:16:34','2020-03-03 06:51:23'),
(10,'Harry Potter and the Deathly Hallows','9781781100134',900,'English',100,340,'hard',3,'https://vignette.wikia.nocookie.net/harrypotter/images/a/ab/Deathly_Hallows_1_poster.jpg',1,640,'2007-07-21','2020-03-03 12:29:04','2020-03-03 12:29:04');

/*Table structure for table `genre` */

DROP TABLE IF EXISTS `genre`;

CREATE TABLE `genre` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `slug` varchar(45) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `slug_UNIQUE` (`slug`,`name`,`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=latin1;

/*Data for the table `genre` */

insert  into `genre`(`id`,`name`,`slug`,`status`,`created_at`,`updated_at`) values 
(1,'History','history',1,'2020-02-27 18:42:01','2020-02-27 17:59:29'),
(2,' Fiction  ','fiction',1,'2020-02-27 18:42:23','2020-02-27 18:42:23'),
(3,' Science  ','science',1,'2020-02-27 18:42:37','2020-02-27 18:42:37'),
(4,' Action  ','action',1,'2020-02-27 18:43:09','2020-02-27 18:43:09'),
(5,' Drama  ','drama',1,'2020-02-27 18:43:16','2020-02-27 18:43:16'),
(6,' Fairytale  ','fairytale',1,'2020-02-27 18:43:22','2020-02-27 18:43:22'),
(7,' Fantasy  ','fantasy',1,'2020-02-27 18:43:34','2020-02-27 18:43:34'),
(8,' Horror  ','horror',1,'2020-02-27 18:43:41','2020-02-27 18:43:41'),
(9,' Mystery  ','mystery',1,'2020-02-27 18:43:47','2020-02-27 18:43:47'),
(10,' Religion  ','religion',1,'2020-02-27 18:43:55','2020-02-27 18:43:55'),
(12,' Thriller  ','thriller',1,'2020-02-27 18:44:06','2020-02-27 18:44:06'),
(13,' Romance  ','romance',1,'2020-02-27 18:44:13','2020-02-27 18:44:13'),
(14,' Poetry  ','poetry',1,'2020-02-27 18:44:18','2020-02-27 18:44:18'),
(15,' Biography  ','biography',1,'2020-02-27 18:44:28','2020-02-27 18:44:28'),
(16,' Autobiography  ','autobiography',1,'2020-02-27 18:44:33','2020-02-27 18:44:33'),
(17,' Travel  ','travel',1,'2020-02-27 18:44:40','2020-02-27 18:44:40'),
(19,' Satire  ','satire',1,'2020-02-27 18:44:52','2020-02-27 18:44:52'),
(25,'Crime','crime',1,'2020-02-27 23:31:25','2020-02-27 18:03:43');

/*Table structure for table `publications` */

DROP TABLE IF EXISTS `publications`;

CREATE TABLE `publications` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `slug` varchar(45) NOT NULL,
  `founding_date` date NOT NULL,
  `description` varchar(300) DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `slug_UNIQUE` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

/*Data for the table `publications` */

insert  into `publications`(`id`,`name`,`slug`,`founding_date`,`description`,`status`,`created_at`,`updated_at`) values 
(1,'Penguin Random House','penguin-random-house','2013-02-20','It\'s considered to be the biggest publishing house in the industry. It has over 200 divisions and imprints',1,'2020-02-28 17:49:43','2020-02-28 17:49:43'),
(2,'Hachette Livre','hachette-livre','1992-07-15','It’s owned by the Lagardère Group and encompasses over 150 imprints. Hachette Livre was officially formed in 1992',1,'2020-02-28 17:50:37','2020-02-28 17:50:37'),
(3,'HarperCollins','harpercollins','1989-11-23','HarperCollins was created in 1989 through a multi-company merger, taking its name from former publishing giants Harper & Row and William Collins',1,'2020-02-28 17:51:10','2020-02-28 17:51:10'),
(4,'Pearson','pearson','2017-12-12','As you might recall, Pearson PLC owns Penguin Random House as well, but its Pearson Education division is limited to academic texts. This is the third of the “big five” educational publishers',1,'2020-02-28 17:51:46','2020-02-28 17:51:46'),
(5,'McGraw-Hill','mcgraw-hill','1960-06-13','McGraw-Hill Education should ring a bell for anyone who’s experienced the magic of the American public school system. As one of the “big five” educational publishers',1,'2020-02-28 17:52:22','2020-02-29 14:55:28');

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `api_key` varchar(60) NOT NULL,
  `status` tinyint(4) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `user` */

insert  into `user`(`id`,`name`,`password`,`api_key`,`status`,`created_at`,`updated_at`) values 
(1,'musbah','49f63510077d7e1b1a386c3a3bfcbd9f','EJfrTQ1bB65AZU18EJX1o4Rx51SU3B',1,'2022-02-09 13:49:33','2022-02-09 13:49:33'),
(2,'abang','49f63510077d7e1b1a386c3a3bfcbd9f','ROl8pNt6ZP8IhB3HVtYJpQiIMMG1zq',1,'2022-02-09 13:56:25','2022-02-09 13:56:25');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
