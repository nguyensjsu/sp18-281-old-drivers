# Startbucks Online System

### Weekly Progress and Challenge
---
```
Week 1: (3/24/18-3/31/18) 
	* Weekly progress: Understanding the final project requirements, analyze and compare different catagories
			in order to find the best topic that can be used and implemented for this final project. 
	* Challenge: Have not started on personal project part, so that it is hard to pick a good topic at this moment.
``` 
```
Week 2: (4/1/18-4/7/18)
        * Weekly progress: 
        * Challenge:
```
```
Week3: (4/8/18-4/15/18)
        * Weekly progress: 
        * Challenge:
```


### 1. Introduction of The Project (04/16/2018)
---
In this project, we decide to create a Starbucks Online Order System. Customers can use this to register, make an online order, check order status, make comments etc. The owner can recevice the order, process it, and also manage the stock, check reviews of each item.


### 2. Architecture Diagram (04/16/2018)
---
This section our team members discussed the roadmap of our project, figure out the tech stack and architecure of the system. As int the figure below, the frontend will be writen by JS, and we use Kong as API Gateway to route API calls to different sub-module. Each module uses Redis as backend database, and Redis is configured in replication mode, at this time, we have not decided how to handle network partition.

### 3. System Design (04/18/2018)
---

#### 3.1 Overview
---
In this project we will build an online Starbucks system. This system contains four main modules

* Order subsystem
* Inventory subsystem
* User management subsystem
* Product review & comments subsystem

Each module is binded with an API server. Each API server interacts with backend database directly, which is [Redis](https://redis.io/) cluster in this project. For simplicity, we use [Kong](https://getkong.org/about/) as an API gateway to route each call to the related API server. We deploy the backend Redis cluster in replication mode. Each quorum has 5 nodes and each module is related to on quorum. [Heroku](https://dashboard.heroku.com/) is used to deploy the application.

#### 3.2 UI Design (04/18/2018)
---

#### 3.3 Kong API Gateway (04/18/2018)
---

#### 3.4 Order Subsystem (04/18/2018)
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
HTTP 204
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

The main challenge here is that, for each write operation, we need to update two different data structures, it may casuse data inconsistance if a failure happns during the operation. Here we use transaction provided by Redis to handle the problem.

#### 3.5 Inventory Subsystem (04/18/2018)
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
/inventory/{id}

Return:
Jsonized inventory struct
```

* Add inventory
```
Method 'POST'
/inventory?name=blacktea&price=10&inventory=100

Return:
Jsonized inventory struct
```

* Update inventory
```
Method 'PUT'
/inventory/id?price=20&inventory=99

Return:
Jsonized inventory struct
```

* Delete inventory
```
Method 'DELETE'
/inventory/{id}

Return:
Status code 204.
```

#### 3.6 User Management Subsystem (04/18/2018)

* Add a new user
```
Method 'POST'
/user?username=email_address&password=123456&firstname=leo&lastname=Peterson&phone=1234567788

Return:
Status code 200
```

* User Authentication
```
```

* Delete user information
```
Method 'DELETE'
/user/{userid}

Return:
Status code 204.
```

#### 3.7 Product Review & Comments Subsystem (04/18/2018)
---

#### 3.8 Error Handling (04/18/2018)
---

### 4. System APIs (04/19/2018)
---

### 5. Kong Configuration (04/19/2018)
---

### 6. Redis Configuration (04/19/2018)
---

### 7. Implementation (04/20/2018)


### 8. Summary
---

### Appendix
---
[Redis](https://redis.io/)
[Redis SDK](https://github.com/go-redis/redis)
[Kong](https://getkong.org/about/)
[Heroku](https://dashboard.heroku.com/)
