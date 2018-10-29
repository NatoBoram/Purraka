-- Items
	-- data-wearableitemid : 2738
	-- data-type : PlayerWearableItem
	-- abstract-icon : /static/img/item/player/icon/c21ea4fb64bd8a37b622af879d45e70d.png
	-- rarity-marker : epic
	-- abstract-name : Petite ceinture à plumes
	-- abstract-type : Ceintures
-- Market
	-- data-itemid : 18342710
	-- currentPrice : 1
	-- buyNowPrice : 1
	-- data-bids : 0

--
-- Database: `eldarya`
--

DROP DATABASE `purraka`;
CREATE DATABASE IF NOT EXISTS `purraka` DEFAULT CHARACTER SET utf8mb4;
USE `purraka`;

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
CREATE TABLE IF NOT EXISTS `items` (
	`data-wearableitemid` int PRIMARY KEY,
	`data-type` varchar(32) not null,
	`abstract-icon` varchar(128) not null,
	`rarity-marker` varchar(16) not null,
	`abstract-name` varchar(64) not null,
	`abstract-type` varchar(32) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `market`
--

DROP TABLE IF EXISTS `market`;
CREATE TABLE IF NOT EXISTS `market` (
	`data-itemid` int PRIMARY KEY,
	`data-wearableitemid` int not null,
	`currentPrice` int not null,
	`buyNowPrice` int not null,
	`data-bids` int not null,
	`active` boolean not null DEFAULT 1,
	CONSTRAINT `fk_market_items`
		FOREIGN KEY (`data-wearableitemid`) REFERENCES items (`data-wearableitemid`)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- View structure for view `market-avgstd`
--

create or replace view `market-avgstd` as  
	select
		`market`.`data-wearableitemid` AS `data-wearableitemid`,
		avg(`market`.`currentPrice`) AS `avg-currentPrice`,
		std(`market`.`currentPrice`) AS `std-currentPrice`,
		avg(`market`.`buyNowPrice`) AS `avg-buyNowPrice`,
		std(`market`.`buyNowPrice`) AS `std-buyNowPrice`,
		avg(`market`.`data-bids`) AS `avg-data-bids`,
		std(`market`.`data-bids`) AS `std-data-bids`
	from `market`
	where
		((`market`.`currentPrice` > 0)
		and (`market`.`buyNowPrice` > 0))
	group by `market`.`data-wearableitemid`
;

-- --------------------------------------------------------

--
-- View structure for view `market-zscore`
--

create or replace view `market-zscore` as
select
	`market`.`data-itemid` AS `data-itemid`,
	coalesce(((`market`.`buyNowPrice` - `market-avgstd`.`avg-buyNowPrice`) / `market-avgstd`.`std-buyNowPrice`),0) AS `zscore-buyNowPrice`,
	coalesce(((`market`.`currentPrice` - `market-avgstd`.`avg-currentPrice`) / `market-avgstd`.`std-currentPrice`),0) AS `zscore-currentPrice`,
	coalesce(((`market`.`data-bids` - `market-avgstd`.`avg-data-bids`) / `market-avgstd`.`std-data-bids`),0) AS `zscore-data-bids`
from `market`
	join `market-avgstd`
where (
	(`market`.`data-wearableitemid` = `market-avgstd`.`data-wearableitemid`)
	and (`market`.`currentPrice` > 0)
	and (`market`.`buyNowPrice` > 0)
	and (`market`.`active` = 1)
);

-- --------------------------------------------------------

--
-- View structure for view `market-everything`
--

create or replace view `market-everything` as
select
	`items`.`data-wearableitemid` AS `data-wearableitemid`,
	`market`.`data-itemid` AS `data-itemid`,
	`items`.`data-type` AS `data-type`,
	`items`.`rarity-marker` AS `rarity-marker`,
	`items`.`abstract-name` AS `abstract-name`,
	`items`.`abstract-type` AS `abstract-type`,
	`market`.`currentPrice` AS `currentPrice`,
	`market-zscore`.`zscore-currentPrice` AS `zscore-currentPrice`,
	`market`.`buyNowPrice` AS `buyNowPrice`,
	`market-zscore`.`zscore-buyNowPrice` AS `zscore-buyNowPrice`,
	`market`.`data-bids` AS `data-bids`,
	`market-zscore`.`zscore-data-bids` AS `zscore-data-bids`,
	`items`.`abstract-icon` AS `abstract-icon`
from `items`
join `market`
join `market-zscore`
where (
	(`items`.`data-wearableitemid` = `market`.`data-wearableitemid`)
	and (`market`.`data-itemid` = `market-zscore`.`data-itemid`)
	and (`market`.`active` = 1)
	and (`market`.`currentPrice` > 0)
	and (`market`.`buyNowPrice` > 0)
);

-- --------------------------------------------------------

--
-- Table structure for table `z-market`
--

CREATE TABLE `z-market` (
  `data-wearableitemid` int(11) not null,
  `data-itemid` int(11) not null,
  `data-type` varchar(32) not null,
  `rarity-marker` varchar(16) not null,
  `abstract-name` varchar(64) not null,
  `abstract-type` varchar(32) not null,
  `currentPrice` int(11) not null,
  `zscore-currentPrice` double not null,
  `buyNowPrice` int(11) not null,
  `zscore-buyNowPrice` double not null,
  `data-bids` int(11) not null,
  `zscore-data-bids` double not null,
  `abstract-icon` varchar(128) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Event structure for event `event_update_zmarket`
--

delimiter |
CREATE EVENT `event_update_zmarket` ON SCHEDULE EVERY 1 MINUTE ENABLE DO begin
	start transaction;
		delete from `z-market`;
		insert into `z-market` select * from `market-everything`;
		commit;
	end |
delimiter ;

-- --------------------------------------------------------

--
-- Table structure for table `sent_on_discord`
--

DROP TABLE IF EXISTS `sent_on_discord`;
CREATE TABLE IF NOT EXISTS `sent_on_discord` (
	`data-itemid` int PRIMARY KEY
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `sent_on_discord`
--

DROP TABLE IF EXISTS `callback_channel`;
CREATE TABLE IF NOT EXISTS `callback_channel` (
	`guild` varchar(32) PRIMARY KEY,
	`channel` varchar(32)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;