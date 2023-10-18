# Restaurant App

The Restaurant Management App is a comprehensive solution for efficiently managing restaurant operations, including order processing, billing, generating monthly income reports, and monitoring menu items and stock levels. Developed using the Go programming language with the Gin web framework and MySQL database.

1. Gin Web Framework
Gin is a lightweight and high-performance web framework for Go. It simplifies routing, request handling, and middleware management, which is ideal for creating a responsive and efficient web-based restaurant management system.

2. MySQL Database
MySQL is a robust and widely-used relational database system. It is well-suited for storing critical restaurant data, such as customer orders, billing information, and menu items. MySQL provides data integrity and reliability, ensuring that the app can handle transactional data effectively.

## Prerequisites

1. Clone this github repository
```sh
git clone https://github.com/prayogatriady/restaurant-management.git
```

2. Adjust `.env` file to connect to your mysql server
3. Insert those tables, the queries are satisfied in migration folder
4. Get golang dependencies and run the program
```sh
cd ..
cd restaurant-management
go mod tidy
go run .
```
5. Download the postman collection in collection folder

## REST API
### for detail, see the postman api collection file
1. First, hit the api `/api/item/genDummyCategories` and `/api/item/genDummyItems` to insert dummy data
2. Hit the api `/api/order/create` to make an orders
3. Hit the api `/api/bill/create` to generate a bill
4. Hit the api `/api/report/create` to generate a monthly report
5. Hit the api `/api/item/itemList` to see a the items menu (stock)