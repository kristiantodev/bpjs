-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 19 Nov 2024 pada 08.20
-- Versi server: 5.7.21-log
-- Versi PHP: 8.0.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bpjs`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `education`
--

CREATE TABLE `education` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `school` varchar(60) NOT NULL,
  `level` varchar(20) NOT NULL,
  `degree` varchar(11) NOT NULL,
  `year_in` int(11) NOT NULL,
  `year_out` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `education`
--

INSERT INTO `education` (`id`, `user_id`, `school`, `level`, `degree`, `year_in`, `year_out`, `created_at`, `updated_at`) VALUES
(2, 4, 'SMP 1 Tangerang', 'SMP', '', 2020, 2024, '2024-11-19 07:09:12', '2024-11-19 07:09:12'),
(3, 4, 'SDN 1 Tangerang', 'SD', '', 2020, 2024, '2024-11-19 07:15:07', '2024-11-19 07:15:07');

-- --------------------------------------------------------

--
-- Struktur dari tabel `skills`
--

CREATE TABLE `skills` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `skill` varchar(60) NOT NULL,
  `level` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `skills`
--

INSERT INTO `skills` (`id`, `user_id`, `skill`, `level`, `created_at`, `updated_at`) VALUES
(1, 4, 'Phyton o', 'Advanced', '2024-11-19 04:48:35', '2024-11-19 04:48:35'),
(3, 4, 'Coding', 'Advanced', '2024-11-19 05:27:39', '2024-11-19 05:27:39'),
(4, 4, 'Phyton', 'Advanced', '2024-11-19 05:27:53', '2024-11-19 05:27:53');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(256) NOT NULL,
  `first_name` varchar(75) NOT NULL,
  `last_name` varchar(75) NOT NULL,
  `gender` varchar(1) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `email` varchar(100) NOT NULL,
  `address` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `first_name`, `last_name`, `gender`, `phone`, `email`, `address`, `created_at`, `updated_at`) VALUES
(1, 'kristianto', 'kris', 'Kristianto', '', 'L', '0899874963', 'kris@gmail.com', 'Tangerang', '2024-10-25 02:09:58', '0000-00-00 00:00:00'),
(2, 'kristianto', 'kris', 'Kristianto', '', 'L', '0899874963', 'kris@gmail.com', 'Tangerang', '2024-10-25 02:11:29', '2024-10-25 02:11:29'),
(3, 'kris', 'pass123', 'Kristianto', '', 'L', '084396432974', 'kris@gmail.com', 'Tangerang', '2024-10-25 02:26:41', '2024-10-25 02:26:41'),
(4, 'bpjs', 'bpjs1', 'Kristianto', '', 'L', '084396432974', 'kris@gmail.com', 'Tangerang', '2024-11-19 03:50:13', '2024-11-19 03:50:13');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `education`
--
ALTER TABLE `education`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `skills`
--
ALTER TABLE `skills`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `education`
--
ALTER TABLE `education`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `skills`
--
ALTER TABLE `skills`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
