# Gook-Items-API

This is one of the [Gook](https://github.com/Armingodiz/Gook) service that is for adding items and searching for them which follow mvc design pattern .

## Features 

* Add item 
* search for items 
* get item with id 

## Dependencies
name     | repo
------------- | -------------
  gorilla-mux | https://github.com/gorilla/mux
  redis       | https://github.com/go-redis/redis
  

## Installation

First make sure you have installed all dependencies ,
make sure you have docker then pull redis image and connect your 8282 port to redis with `sudo docker run --name redis-usdb -p 8282:6379 -d redis`.
Then just simply clone this repository and start service with `go run main.go` (your service will be running on `127.0.0.1:3333`)



## EndPoints 

	POST ==> /items  (Create item)
	GET ==> /items/{id} (Get item by id)
	POST ==> /items/search (Search for users)


## TODO 
* Delete item
* Update item
* Sell item 
