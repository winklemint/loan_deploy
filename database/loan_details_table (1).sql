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
-- Table structure for table `loan_details_table`
--

CREATE TABLE `loan_details_table` (
  `id` int(11) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `applicant_contact` varchar(10) DEFAULT NULL,
  `loan_type` varchar(40) DEFAULT NULL,
  `loan_amount` float DEFAULT NULL,
  `pincode` int(11) DEFAULT NULL,
  `tenure` int(11) DEFAULT NULL,
  `employment_type` varchar(90) DEFAULT NULL,
  `gross_monthly_income` float DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `last_modified` datetime DEFAULT current_timestamp(),
  `status` varchar(45) DEFAULT 'Pending',
  `is_delete` int(11) DEFAULT 0,
  `user_id` int(11) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin_name` varchar(255) DEFAULT '123',
  `progress_status` varchar(255) NOT NULL DEFAULT 'None',
  `enq_moved_to_lead` tinyint(1) NOT NULL DEFAULT 0,
  `device_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `loan_details_table`
--

INSERT INTO `loan_details_table` (`id`, `full_name`, `applicant_contact`, `loan_type`, `loan_amount`, `pincode`, `tenure`, `employment_type`, `gross_monthly_income`, `created_at`, `last_modified`, `status`, `is_delete`, `user_id`, `remark`, `admin_name`, `progress_status`, `enq_moved_to_lead`, `device_id`) VALUES
(1, 'Swati Chouhan', '7987436119', 'Home', 1200000, 480001, 2, 'BussinessOwner', 40000, '2023-08-19 17:50:53', '2023-08-19 17:50:53', 'Pending', 1, 1, NULL, '123', 'None', 1, '4b2325d9-0241-497c-9a59-6a974f888208'),
(2, '', NULL, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 12:25:42', '2023-08-22 12:25:42', 'Pending', 0, 1, NULL, '123', 'None', 0, NULL),
(3, '', NULL, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 12:25:53', '2023-08-22 12:25:53', 'Pending', 0, 1, NULL, '123', 'None', 0, NULL),
(4, '', NULL, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 12:26:07', '2023-08-22 12:26:07', 'Pending', 0, 1, NULL, '123', 'None', 0, NULL),
(5, '', NULL, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 12:26:10', '2023-08-22 12:26:10', 'Pending', 0, 1, NULL, '123', 'None', 0, NULL),
(6, '', NULL, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 12:26:12', '2023-08-22 12:26:12', 'Pending', 0, 1, NULL, '123', 'None', 0, NULL),
(7, '', NULL, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 12:26:13', '2023-08-22 12:26:13', 'Pending', 0, 1, NULL, '123', 'None', 0, NULL),
(8, '', NULL, 'car', 90000, 480001, 1, 'self', 30000, '2023-08-22 12:26:14', '2023-08-22 12:26:14', 'Pending', 0, 1, NULL, '123', 'None', 0, NULL),
(9, '', NULL, 'BuyHome', 0, 0, 0, 'salaried', 0, '2023-08-23 13:36:14', '2023-08-23 13:36:14', 'Pending', 0, NULL, NULL, '123', 'None', 0, '4b2325d9-0241-497c-9a59-6a974f888208'),
(10, 'Swati Chouhan', '7987436119', 'Home', 0, 0, 0, '', 0, '2023-08-24 15:50:33', '2023-08-24 15:50:33', 'Pending', 0, NULL, NULL, '123', 'None', 0, '4b2325d9-0241-497c-9a59-6a974f888208; aplicant_contact=7987436119; full_name=Swati%20Chouhan; applicant_contact=7987436119; loan_type=Home'),
(11, 'Swati Chouhan', '7987436119', 'Home', 0, 0, 0, '', 0, '2023-08-24 15:52:04', '2023-08-24 15:52:04', 'Pending', 0, NULL, NULL, '123', 'None', 0, '4b2325d9-0241-497c-9a59-6a974f888208; aplicant_contact=7987436119; full_name=Swati%20Chouhan; applicant_contact=7987436119; loan_type=Home; loan_id=10'),
(12, 'Swati Chouhan', '7987436119', 'Home', 0, 0, 0, '', 0, '2023-08-24 15:55:54', '2023-08-24 15:55:54', 'Pending', 0, NULL, NULL, '123', 'None', 0, '4b2325d9-0241-497c-9a59-6a974f888208; aplicant_contact=7987436119; full_name=Swati%20Chouhan; applicant_contact=7987436119; loan_type=Home; loan_id=11'),
(13, 'Swati Chouhan', '7987436119', 'Home', 0, 0, 0, '', 0, '2023-08-24 16:00:47', '2023-08-24 16:00:47', 'Pending', 0, NULL, NULL, '123', 'None', 0, '4b2325d9-0241-497c-9a59-6a974f888208; aplicant_contact=7987436119; full_name=Swati%20Chouhan; applicant_contact=7987436119; loan_type=Home; loan_id=12'),
(14, 'Swati', '7987436188', 'Home', 0, 0, 0, '', 0, '2023-08-24 17:49:08', '2023-08-24 17:49:08', 'Pending', 0, NULL, NULL, '123', 'None', 0, '84cb33f1-e18f-4d05-b6b1-bf4ed7c1b178'),
(15, 'Swati', '7987436188', 'Home', 0, 0, 0, '', 0, '2023-08-24 17:49:22', '2023-08-24 17:49:22', 'Pending', 0, NULL, NULL, '123', 'None', 0, '84cb33f1-e18f-4d05-b6b1-bf4ed7c1b178'),
(16, 'Swati', '7987436188', 'Home', 0, 0, 0, '', 0, '2023-08-24 17:55:10', '2023-08-24 17:55:10', 'Pending', 0, NULL, NULL, '123', 'None', 0, '84cb33f1-e18f-4d05-b6b1-bf4ed7c1b178'),
(17, 'Swati', '7987436188', 'Home', 0, 0, 0, '', 0, '2023-08-24 17:56:46', '2023-08-24 17:56:46', 'Pending', 0, NULL, NULL, '123', 'None', 0, '84cb33f1-e18f-4d05-b6b1-bf4ed7c1b178'),
(18, 'Swati', '7987436188', 'Home', 0, 0, 0, '', 0, '2023-08-24 18:01:16', '2023-08-24 18:01:16', 'Pending', 0, NULL, NULL, '123', 'None', 0, '84cb33f1-e18f-4d05-b6b1-bf4ed7c1b178'),
(19, 'Swati', '7987436188', 'Home', 0, 0, 0, '', 0, '2023-08-24 18:02:07', '2023-08-24 18:02:07', 'Pending', 0, NULL, NULL, '123', 'None', 0, '84cb33f1-e18f-4d05-b6b1-bf4ed7c1b178'),
(20, 'Swati Chouhan', '7987436119', 'Personal', 0, 0, 0, '', 0, '2023-08-24 18:10:27', '2023-08-24 18:10:27', 'Pending', 0, NULL, NULL, '123', 'None', 0, NULL),
(21, 'Swati Chouhan', '7987436119', 'wheeler', 0, 0, 0, '', 0, '2023-08-24 19:03:12', '2023-08-24 19:03:12', 'Pending', 0, NULL, NULL, '123', 'None', 0, 'cf110ad5-1b29-4c3e-ad07-f9ae5921f730'),
(22, 'Swati Chouhan', '7987436119', 'Gold', 0, 0, 0, '', 0, '2023-08-24 20:21:24', '2023-08-24 20:21:24', 'Pending', 0, NULL, NULL, '123', 'None', 0, '37376eca-d7b0-4690-815e-c68e71547931');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `loan_details_table`
--
ALTER TABLE `loan_details_table`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `loan_details_table`
--
ALTER TABLE `loan_details_table`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
