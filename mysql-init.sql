CREATE DATABASE IF NOT EXISTS auction_db;

-- Switch to the 'test' database
USE auction_db;

CREATE TABLE IF NOT EXISTS ad_spaces (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    base_price FLOAT NOT NULL,
    end_time DATETIME NOT NULL,
    current_bid FLOAT DEFAULT 0,
    winner_id INT DEFAULT 0
);


CREATE TABLE IF NOT EXISTS bids (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ad_space_id INT NOT NULL,
    bidder_id INT NOT NULL,
    bid_amount FLOAT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ad_space_id) REFERENCES ad_spaces(id)
);

CREATE TABLE IF NOT EXISTS bidders (
   id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   name varchar(255) DEFAULT NULL,
   email varchar(255) DEFAULT NULL
)