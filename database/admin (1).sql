-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 25, 2023 at 12:31 PM
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
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `id` int(11) NOT NULL,
  `username` varchar(45) NOT NULL,
  `password` varchar(255) NOT NULL,
  `contact` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id`, `username`, `password`, `contact`, `email`, `role`) VALUES
(19, 'Raksha Rajput', '2432612431302474584a4572592e686c6d644a464b335343677542492e79674c552e62306e6d6e65615a72664d2f7a64534b33547a67533235357471', 434354343, 'raksha.rajput98@gmail.com', 'Super Admin'),
(25, 'pooja', '243261243130244e663052526b35633148725679324a72556b4652306561467a2f46523032697069724b756d6b5a6442584b5676526677786d4c2e57', 1234567890, 'pooja@gmail.com', 'Sub Admin'),
(123, 'admin_name', '24326124313024564637682e31422e597479392f7359694c47502e4b2e4438586c61616e2f65516f745977304170342f4965686b6e77654d6974496d', 1234567890, 'hjkshjdk@gmail.com', 'Super Admin'),
(125, 'Swati Chouhan', '24326124313024766f776c663856534e2e746a496a5844686662416c65356d5075732f5a6c2f385a7532695a69497941345053624959335468786936', 2147483647, 'swatichouhan39@gmail.com', 'Sub Admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=126;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
