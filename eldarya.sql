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

DROP DATABASE `eldarya`;
CREATE DATABASE IF NOT EXISTS `eldarya` DEFAULT CHARACTER SET utf8;
USE `eldarya`;

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
CREATE TABLE IF NOT EXISTS `items` (
	`data-wearableitemid` int PRIMARY KEY,
	`data-type` varchar(32) DEFAULT NULL,
	`abstract-icon` varchar(128) DEFAULT NULL,
	`rarity-marker` varchar(16) DEFAULT NULL,
	`abstract-name` varchar(32) DEFAULT NULL,
	`abstract-type` varchar(16) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `market`
--

DROP TABLE IF EXISTS `market`;
CREATE TABLE IF NOT EXISTS `market` (
	`data-itemid` int PRIMARY KEY,
	`data-wearableitemid` int NOT NULL,
	`currentPrice` int DEFAULT NULL,
	`buyNowPrice` int DEFAULT NULL,
	`data-bids` int DEFAULT NULL,
	`active` boolean NOT NULL DEFAULT 1,
	CONSTRAINT `fk_market_items`
		FOREIGN KEY (`data-wearableitemid`) REFERENCES items (`data-wearableitemid`)
		ON DELETE CASCADE
		ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- View structure for view `market-view`
--

create view `market-average` as  
	select
		`data-wearableitemid`,
		avg(`currentPrice`) as `average-currentPrice`,
		avg(`buyNowPrice`) as `average-buyNowPrice`,
		avg(`data-bids`) as `average-data-bids`
	from `market`
	group by `data-wearableitemid`
;

-- Difference between current values and averages

create view `market-diff` as  
select
	`data-itemid`,
	(`currentPrice` - (select `average-currentPrice` from `market-average` where `market`.`data-wearableitemid` = `market-average`.`data-wearableitemid`)) as `diff-currentPrice`,
	(`buyNowPrice` - (select `average-buyNowPrice` from `market-average` where `market`.`data-wearableitemid` = `market-average`.`data-wearableitemid`)) as `diff-buyNowPrice`,
	(`data-bids` - (select `average-data-bids` from `market-average` where `market`.`data-wearableitemid` = `market-average`.`data-wearableitemid`)) as `diff-data-bids`
from `market`;

-- Average of differences

create view `market-sigma` as 
select
	`data-wearableitemid`,
	avg(abs(`diff-currentPrice`)) as `sigma-currentPrice`,
	avg(abs(`diff-buyNowPrice`)) as `sigma-buyNowPrice`,
	avg(abs(`diff-data-bids`)) as `sigma-data-bids`
from `market`, `market-diff`
where `market`.`data-itemid` = `market-diff`.`data-itemid`
group by `data-wearableitemid`;

-- Z Score

create view `market-zscore` as
select
	`market`.`data-itemid`,
	(`diff-currentPrice` / `sigma-currentPrice`) as `zscore-currentPrice`,
	(`diff-buyNowPrice` / `sigma-buyNowPrice`) as `zscore-buyNowPrice`,
	(`diff-data-bids` / `sigma-data-bids`) as `zscore-data-bids`
from `market-diff`, `market-sigma`, `market`
where `market`.`data-itemid` = `market-diff`.`data-itemid`
	and `market`.`data-wearableitemid` = `market-sigma`.`data-wearableitemid`
;