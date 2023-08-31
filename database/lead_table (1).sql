-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 25, 2023 at 12:32 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `loan_application`
--

-- --------------------------------------------------------

--
-- Table structure for table `lead_table`
--

CREATE TABLE `lead_table` (
  `id` int(11) NOT NULL,
  `loan_type` varchar(45) DEFAULT NULL,
  `loan_amount` int(11) DEFAULT NULL,
  `pincode` int(11) DEFAULT NULL,
  `tenure` int(11) DEFAULT NULL,
  `employment_type` varchar(45) DEFAULT NULL,
  `gross_monthly_income` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `last_modified` timestamp NOT NULL DEFAULT current_timestamp(),
  `status` varchar(9) NOT NULL DEFAULT 'Pending',
  `is_delete` int(11) DEFAULT 0,
  `user_id` int(11) NOT NULL,
  `remark` varchar(255) NOT NULL DEFAULT 'null',
  `admin_id` varchar(255) NOT NULL DEFAULT '123',
  `progress_status` varchar(255) NOT NULL DEFAULT 'None',
  `enq_moved_to_lead` tinyint(1) NOT NULL,
  `device_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `lead_table`
--

INSERT INTO `lead_table` (`id`, `loan_type`, `loan_amount`, `pincode`, `tenure`, `employment_type`, `gross_monthly_income`, `created_at`, `last_modified`, `status`, `is_delete`, `user_id`, `remark`, `admin_id`, `progress_status`, `enq_moved_to_lead`, `device_id`) VALUES
(1, 'BuyHome', 1200000, 480001, 2, 'BussinessOwner', 40000, '2023-08-19 12:30:45', '2023-08-19 12:20:53', 'Approved', 1, 1, 'ICICI bank := approved\nHDFC bank := declined', '19', 'Completed', 1, NULL),
(2, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:28', '2023-08-22 06:57:28', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(3, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:35', '2023-08-22 06:57:35', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(4, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:36', '2023-08-22 06:57:36', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(5, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:37', '2023-08-22 06:57:37', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(6, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:37', '2023-08-22 06:57:37', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(7, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:37', '2023-08-22 06:57:37', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(8, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:38', '2023-08-22 06:57:38', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(9, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 06:57:39', '2023-08-22 06:57:39', 'Pending', 0, 1, 'null', '123', 'None', 0, NULL),
(10, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 10:10:26', '2023-08-22 06:57:40', 'Pending', 0, 1, 'gold', '19', 'None', 0, NULL),
(11, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 10:08:41', '2023-08-22 06:57:41', 'Approved', 0, 1, 'car car', '19', 'In Progress', 0, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `lead_table`
--
ALTER TABLE `lead_table`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `lead_table`
--
ALTER TABLE `lead_table`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
