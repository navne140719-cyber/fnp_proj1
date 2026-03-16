# Order Processing System (Kafka + Golang + Spring Boot)

## GitHub Repository Link
https://github.com/navne140719-cyber/fnp_proj1

---

## Steps to Run the Project

### 1. Start Required Services

Run Docker containers:

docker compose up -d

Verify containers:

docker ps

---

### 2. Create Kafka Topic

docker exec -it kafka1 kafka-topics \
--create \
--topic order-created \
--bootstrap-server localhost:9092 \
--partitions 1 \
--replication-factor 1

---

### 3. Setup Database

Create database:

CREATE DATABASE order_system;
USE order_system;

Create tables:

users  
products  
orders  
order_items

Insert sample data:

INSERT INTO users(name,email) VALUES ('rohan','rohan@gmail.com');

INSERT INTO products(name,price,stock) VALUES
('Laptop',23000,20),
('Mouse',500,50),
('Keyboard',1200,30);

---

### 4. Run Go Order Service

cd go-service
cd cmd

go run main.go

---

### 5. Run Spring Boot Consumer

Open spring-service in IntelliJ and run:

OrderConsumerApplication.java

or

mvn spring-boot:run

---

### 6. Test API

POST http://localhost:8080/orders

Request body:

{
 "userId": 1,
 "items": [
   {
     "productId": 1,
     "qty": 2
   }
 ]
}
