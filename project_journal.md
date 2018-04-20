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

#### 3.5 Inventory Subsystem
---
* Get inventory

    https://localhost:8000/inventory/id

    Method 'GET'

* Add inventory

    https://localhost:8000/inventory?name=blacktea&price=10&inventory=100

    Method 'POST'

* Update inventory

    https://localhost:8000/inventory/id?price=20&inventory=99

    Method 'PUT'

* Delete inventory

    https://localhost:8000/inventory/id

    Method 'DELETE'

#### 3.6 User Management Subsystem
---

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