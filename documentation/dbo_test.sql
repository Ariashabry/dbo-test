-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 21, 2024 at 11:40 PM
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
-- Database: `dbo_test`
--

-- --------------------------------------------------------

--
-- Table structure for table `history_cart`
--

CREATE TABLE `history_cart` (
  `id_cart` bigint(20) UNSIGNED NOT NULL,
  `id_order` bigint(20) UNSIGNED DEFAULT NULL,
  `id_user` bigint(20) UNSIGNED DEFAULT NULL,
  `id_product` bigint(20) UNSIGNED DEFAULT NULL,
  `price` double DEFAULT NULL,
  `qty` bigint(20) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `history_cart`
--

INSERT INTO `history_cart` (`id_cart`, `id_order`, `id_user`, `id_product`, `price`, `qty`, `created_at`) VALUES
(1, 2, 8, 1, 153500, 1, '2024-01-21 22:47:33'),
(2, 2, 8, 2, 5500, 2, '2024-01-21 22:47:33'),
(3, 3, 7, 1, 153500, 1, '2024-01-22 01:37:24'),
(4, 3, 7, 2, 5500, 2, '2024-01-22 01:37:24'),
(9, 6, 7, 1, 153500, 1, '2024-01-22 01:48:44'),
(10, 6, 7, 2, 5500, 2, '2024-01-22 01:48:44');

-- --------------------------------------------------------

--
-- Table structure for table `order`
--

CREATE TABLE `order` (
  `id_order` bigint(20) UNSIGNED NOT NULL,
  `id_user` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `total_amount` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `order`
--

INSERT INTO `order` (`id_order`, `id_user`, `created_at`, `updated_at`, `total_amount`) VALUES
(2, 8, '2024-01-21 22:47:33', '2024-01-21 22:47:33', 164500),
(3, 7, '2024-01-22 01:37:24', '2024-01-22 01:37:24', 164500),
(6, 8, '2024-01-22 01:48:44', '2024-01-22 01:48:44', 164500);

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id_product` bigint(20) UNSIGNED NOT NULL,
  `product_name` varchar(100) DEFAULT NULL,
  `brand` varchar(100) DEFAULT NULL,
  `price` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id_product`, `product_name`, `brand`, `price`) VALUES
(1, 'Cat Tembok Catylac 5 Kg', 'Dulux', 153500),
(2, 'Cat anti bocor No drop Waterproofing 1 kg', 'No Drop', 55000),
(3, 'Kuas Cat Tembok YSK Ukuran 4 Inch Type 633 Bulu Putih Kayu Vernis', 'YSK', 6000);

-- --------------------------------------------------------

--
-- Table structure for table `role`
--

CREATE TABLE `role` (
  `id_role` bigint(20) UNSIGNED NOT NULL,
  `role` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `role`
--

INSERT INTO `role` (`id_role`, `role`) VALUES
(1, 'customer');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id_user` bigint(20) UNSIGNED NOT NULL,
  `full_name` varchar(50) DEFAULT NULL,
  `gender` varchar(20) DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `user_role` bigint(20) UNSIGNED DEFAULT NULL,
  `user_name` varchar(50) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id_user`, `full_name`, `gender`, `birth_date`, `created_at`, `updated_at`, `user_role`, `user_name`, `password`) VALUES
(7, 'Prabowo S', 'Pria', '1977-12-11', '2024-01-20 23:57:03', '2024-01-21 13:56:12', 1, 'prabowo', '$2a$10$6y8NIHi28I74Hj8NEEH/kuUd9kx/Q5dzCvniq6I265TxckLCpxNLi'),
(8, 'Anies Baswedan', 'Pria', '1981-11-25', '2024-01-21 13:54:26', '2024-01-21 13:54:26', 1, 'anies', '$2a$10$a1MPpH8afP23qDpumapDveP.TJxpfSB1amvEtrePLWEpSKA9YH3BS'),
(9, '', '', '0000-00-00', '2024-01-22 02:57:03', '2024-01-22 02:57:03', 1, 'shabry123', '$2a$10$W2Zwy7jXTXYbyemva9qM9eYzFq1hXOM9eJSOeUZPTTdHcovvBgPR.'),
(10, 'Cak Imin', 'Pria', '1981-11-25', '2024-01-22 03:35:42', '2024-01-22 03:43:50', 1, 'muhaimin', '$2a$10$g6jFpgirx3856OnwUkWuwuOpR/3X8d3ElMWrBcgKwTVoygD14OPzy'),
(11, 'Cak Imin', 'Pria', '1981-11-25', '2024-01-22 03:44:50', '2024-01-22 03:44:50', 1, 'muhaimin', '$2a$10$8Qb5lRU4toMgzoHzKSm5hOUdGssktO.mNB4IIoWc.tbBIlHm98nua');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `history_cart`
--
ALTER TABLE `history_cart`
  ADD PRIMARY KEY (`id_cart`),
  ADD KEY `fk_history_cart_product` (`id_product`),
  ADD KEY `fk_order_cart` (`id_order`);

--
-- Indexes for table `order`
--
ALTER TABLE `order`
  ADD PRIMARY KEY (`id_order`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id_product`);

--
-- Indexes for table `role`
--
ALTER TABLE `role`
  ADD PRIMARY KEY (`id_role`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`) USING BTREE,
  ADD KEY `fk_user_role` (`user_role`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `history_cart`
--
ALTER TABLE `history_cart`
  MODIFY `id_cart` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `order`
--
ALTER TABLE `order`
  MODIFY `id_order` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `product`
--
ALTER TABLE `product`
  MODIFY `id_product` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `role`
--
ALTER TABLE `role`
  MODIFY `id_role` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id_user` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `history_cart`
--
ALTER TABLE `history_cart`
  ADD CONSTRAINT `fk_history_cart_product` FOREIGN KEY (`id_product`) REFERENCES `product` (`id_product`),
  ADD CONSTRAINT `fk_order_cart` FOREIGN KEY (`id_order`) REFERENCES `order` (`id_order`);

--
-- Constraints for table `user`
--
ALTER TABLE `user`
  ADD CONSTRAINT `fk_user_role` FOREIGN KEY (`user_role`) REFERENCES `role` (`id_role`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
