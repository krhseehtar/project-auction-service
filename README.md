# project-auction-service
Auction service for ad space on a website

This repository contains a Go-based auction service for managing ad spaces on a website. The service is divided into two modules: the Supply Side Service and the Demand Side Service, responsible for listing available ad spaces and managing bidder activities, respectively.

## Supply Side Service

## API Endpoints
    POST supply-service/adspaces
    GET supply-service/adspaces
    GET supply-service/adspaces/{id}
    GET supply-service/adspaces/{id}/winner

### AdSpace Data Model
    ID: Unique identifier for the ad space.
    Name: Name of the ad space.
    Base Price: Initial price of the ad space.
    EndTime: Date and time when the auction for the ad space ends.
    CurrentBid: Current highest bid for the ad space.
    WinnerID: ID of the winning bidder (if auction is completed).
### Constraints/Assumptions
    Ad spaces cannot have negative base prices.
    Auction end time must be in the future.


## Demand Side Service

### API Endpoints
    POST /demand_service/bidders
    GET /demand_service/bidders
    GET /demand_service/bidders/{id}
    POST demand_service/bidders/{id}/bids
    GET demand_service/adspaces/{id}/bids
    GET demand_service/bidders/{id}/bids
    GET demand_service/bidders/{id}/adspaces/{adspaceID}/bids

### Bidder Data Model
    ID: Unique identifier for the bidder.
    Name: Name of the bidder.
    Email: Email address of the bidder.

### Bid Data Model
    ID: Unique identifier for the bid.
    AdSpaceID: ID of the ad space for which the bid is placed.
    BidderID: ID of the bidder placing the bid.
    BidAmount: Amount of the bid.
    Timestamp: Date and time when the bid is placed.

### Constraints/Assumptions
    Bids cannot be negative.
    Bids amount cannot be less than or equal to current bid
    Bidders must register before placing a bid.
    Bids must be placed before the end time of an auction.

## Auction Process
    The auction starts with the ad space being created at its base price.
    Bidders must register before placing a bid.
    Bidders place bids higher than the current highest bid.
    Bidding continues until the auction end time is reached.
    The bidder with the highest bid at the end of the auction wins the ad space.


## MySQL Database Schema
    The MySQL database includes tables for ad_spaces, bidders, and bids to store relevant data. 
    Foreign key constraints are defined between tables where necessary.
    
    ad_spaces {id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR, base_price FLOAT, end_time DATETIME,
    current_bid FLOAT, winner_id INT}

    bids {id INT AUTO_INCREMENT PRIMARY KEY, ad_space_id INT, bidder_id ,bid_amount FLOAT, timestamp TIMESTAMP,
    FOREIGN KEY (ad_space_id) REFERENCES ad_spaces(id)}

    bidders {id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar, email varchar}


## Installation
    To run this application, you need to have Go and Docker installed on your machine.
    Clone this repository: git clone https://github.com/krhseehtar/project-auction-service.git
    Change into the directory: cd project-auction-service/
    Start the application: docker-compose up --build


## Usage
    Once the application is running, you can supply-services at http://localhost:8080/supply-service
    and demand-services http://localhost:8081/demand-service. 

    
## API Collection
https://api.postman.com/collections/10344094-ec0014fc-64a7-4a75-9295-69b6723b1b07?access_key=PMAT-01HBTBFQMWH65CRWCNWR9PRMB1


