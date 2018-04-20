# Startbucks Online System

### 1. Introduction of The Project
---
In this project we will build an online Starbucks system. This system contains four main modules
* Order subsystem
* Inventory subsystem
* User management subsystem
* Product review & comments subsystem

### 2. Architecture Diagram
---

### 3. System Design
---

#### 3.1 Overview
---
This system contains 4 modules, each module is binded with an API server. Each API server interacts with backend database directly, which is [Redis](https://redis.io/) cluster in this project. For simplicity, we use [Kong](https://getkong.org/about/) as an API gateway to route each call to the related API server. We deploy the backend Redis cluster in replication mode. Each quorum has 5 nodes and each module is related to on quorum. [Heroku](https://dashboard.heroku.com/) is used to deploy the application.

#### 3.2 UI Design
---

#### 3.3 Kong API Gateway
---

#### 3.4 Order Subsystem
---

Data Structure
```
struct Order {
	string orderId;
	string userId;
	List<string> items;
	int status; // 0 for created, 1 for completed, 2 for failed
}
```

Database Schema
```
1. Use Redis SET APIs to store order info
   Key is order id, and value is a jsonized order struct

2. Use Redis LIST APIs to storge users' order info
   Key is user id, and value is all the orders the user made
```

APIs

* Create Order
```
Method: POST
/order?userid=xxx&items=xxx,yyy,zzz

Response:
Jsonized order structure
```

* Update Order
```
Method: POST
/order/{orderid}?add=xx,yy&delete=zz

Response:
Jsonized order structure
```

* Delete Order
```
Method: DELETE
/order/{orderid}

Response:
OK status
```

* Get Order
```
Method: GET
/order/{orderid}

Response:
Jsonized Order struct or 404 error code if not exist
```

* Get user orders
```
Method: GET
/order?userid=xxx

Response:
All orders belone to the user
```

#### 3.5 Inventory Subsystem
---
Data Stucture
```
struct Inventory {
    String inventoryId;
    String inventoryName;
    Double inventoryPrice;
    Integer inventoryLeft;
}
```

APIs

* Get inventory
```
Method 'GET'
https://localhost:8000/inventory/id

Return:
Jsonized inventory struct
```

* Add inventory
```
Method 'POST'
https://localhost:8000/inventory?name=blacktea&price=10&inventory=100

Return:
Jsonized inventory struct
```

* Update inventory
```
Method 'PUT'
https://localhost:8000/inventory/id?price=20&inventory=99

Return:
Jsonized inventory struct
```

* Delete inventory
```
Method 'DELETE'
https://localhost:8000/inventory/id

Return:
Status code 200.
```

#### 3.6 User Management Subsystem

* Add a new user
```
Method 'POST'
https://localhost:8000/user?username=email_address&password=123456&firstname=leo&lastname=Peterson&phone=1234567788

Return:
Status code 200
```

* User Authentication
```
```

* Delete user information
```
Method 'DELETE'
https://localhost:8000/user/username

Return:
Status code 200.
```

#### 3.7 Product Review & Comments Subsystem
---

#### 3.8 Error Handling
---

### 4. System APIs
---

### 5. Kong Configuration
---

### 6. Redis Configuration
---

### 7. Summary
---
