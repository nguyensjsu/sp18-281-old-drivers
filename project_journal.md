# Startbucks Online System

### Weekly Progress and Challenge
---
```
Week 1: (3/24/18-3/31/18) 
	* Weekly progress: Understanding the final project requirements, analyze and compare different 
                       catagories in order to find the best topic that can be used and implemented 
                       for this final project. 

	* Challenge: Have not started on personal project part, so that it is hard to pick a good 
                 topic at this moment.
``` 
```
Week 2: (4/1/18-4/7/18)
    * Weekly progress: Team members are still working on individual project part, such as Redis DB 
                       installtion and configuration part on EC2 instances.

    * Challenge: Testing on AP.
```
```
Week3: (4/7/18-4/14/18)
    * Weekly progress: Discuss on how to combine individual's pieces on team project, and the 
                       selection on deployment mode for Redis and network routing.

    * Challenge: Combine individual's pieces on team project, and the selection on deployment mode 
                 for Redis and network routing.
```
```
Week4: (4/14/18-4/21/18)
    * Weekly progress: Starting working on architecture of team project. We decide to work on coffee 
                       order system. You can see more detail on each section in below. 

    * Challenge: Implementation, Testing part on each API.
```
```
Week5: (4/21/18-4/28/18)
    * Weekly progress: Continues working on implementation part of API in different catagories for 
                       all team members, and testing done for API parts for all team members.
                       Finish Kong configuration on AWS.

    * Challenge: Implementation, Testing part on each API.
```
```
Week6: (4/28/18-5/4/18)
    * Weekly progress: Enhance GO API server, supporting data sharding, replication and consistency.
    * Challenge: Support isolation accross data shards, and maintain data consistency
```


### 1. Introduction of The Project (04/16/2018)
---
In this project, we decide to create a Starbucks Online Order System. Customers can use this to 
register, make an online order, check order status, make comments etc. The owner can recevice 
the order, process it, and also manage the stock, check reviews of each item.


### 2. Architecture Diagram (04/16/2018)
---
This section our team members discussed the roadmap of our project, figure out the tech stack and 
architecure of the system. As int the figure below, the frontend will be writen by JS, and we use 
Kong as API Gateway to route API calls to different sub-module. Each module uses Redis as backend 
database, and Redis is configured in replication mode, at this time, we have not decided how to 
handle network partition.

### 3. System Design (04/18/2018)
---

#### 3.1 Overview
---
In this project we will build an online Starbucks system. This system contains four main modules

* Order subsystem
* Inventory subsystem
* User management subsystem
* Product review & comments subsystem

Each module is binded with an API server. Each API server interacts with backend database directly, 
which is [Redis](https://redis.io/) cluster in this project. For simplicity, we use [Kong](https://getkong.org/about/) 
as an API gateway to route each call to the related API server. We deploy the backend Redis cluster 
in replication mode. Each quorum has 5 nodes and each module is related to on quorum. [Heroku](https://dashboard.heroku.com/) 
is used to deploy the application.

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
	Date date;
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
/users/{userid}/order?&items=xxx,yyy,zzz

Response:
Jsonized order structure
```

* Update Order
```
Method: POST
/users/{userid}/order/{orderid}?add=xx,yy&delete=zz

Response:
Jsonized order structure
```

* Delete Order
```
Method: DELETE
/users/{userid}/order/{orderid}

Response:
HTTP 204
```

* Get Order
```
Method: GET
/users/{userid}/order/{orderid}

Response:
Jsonized Order struct or 404 error code if not exist
```

* Get user orders
```
Method: GET
/users/{userid}/orders

Response:
All orders belone to the user
```

The main challenge here is that, for each write operation, we need to update two 
different data structures, it may casuse data inconsistance if a failure happns 
during the operation. Here we use transaction provided by Redis to handle the problem.

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
/inventory?name=blacktea&price=10&amount=100

Return:
Jsonized inventory struct
```

* Update inventory
```
Method 'PUT'
/inventory/id?price=20&amount=99


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
---
Data structure
```
type User struct {
    UserId string
    UserName string
    Phone string
    Balance int
}
```

APIs

* Add a new user
```
Method 'POST'
/user?name=Leo&phone=1234567788&balance=100

Return:
Jsonized user struct
```

* Get user
```
Method 'GET'
/user/{id}

Return:
Jsonized user struct
```

* Update user
```
Method 'PUT'
/user/id?phone=1234567789&balance=150

Return:
Jsonized user struct
```

* Delete user information
```
Method 'DELETE'
/user/{id}

Return:
Status code 204.
```

#### 3.7 Product Review & Comments Subsystem (04/18/2018)
---
* Data Structure   
```
type Review struct {
  ReviewId string
  UserId  string
  Item    string
  Content string
  Date    string
}
```

* Add a review
```
Method 'POST'

/users/{userid}/review?content=xxx

Return:
Jsonized review struct
```

* Get review
```
Method 'GET'

/users/{userid}/review/{reviewid}

Return:
Jsonized review struct
```

* Update review
```
Method 'POST'

/users/{userid}/review/{reviewid}

Jsonized review struct
```

* Delete review
```
Method 'DELETE'

/users/{userid}/review/{reviewid}


Return:
Status code 204.
```

* Get review by user
```
Method 'GET'

/users/{userid}/reviews

Response:
All reviews belong to the user
```

### 4. AWS Configuration (01/05/2018)
---
* Create Instance
Each module will do the same operations.
```
Create 6 instances on AWS EC2 with designated VPC (vpc-7efd6019). Two of them in on subnet (subnet-19314642), another four in subnet (subnet-f7c59e90).
```

* Create Network ACL
```
The ACL rules for subnet-19314642 allow all the network protocols except two ports 6379 on TCP. This is because 
redis servers use the port to communicate with each other through TCP. After 
applying the ACL rules, subnet-19314642 will act as been partitionned, because redis nodes in subnet1 can only communicate 
with each other but can't talk to nodes in other subnets.
```
The purpose of creating customized network ACL and subnets is for simulating network partition. This process can verify our system's capability to handle network partition, and demonstrate the advantage of our sharding solution, that is even part of data is unavailable, our system can still work properly for the rest of the dataset.

### 5. Redis Configuration (04/19/2018)
---
We will create [Redis cluster](https://redis.io/topics/cluster-tutorial) for our service. Each module has its own redis cluster. For each cluster, the key space is split into 16384 hash slots. So data is sharded on the 16384 slots according to the key value. To ensure the transactional operation
on some API, we use hash tag in Redis to make sure all the keys with the same hash tag are dispatched to the same slot.

* Initialization
```
# ips is a file with all the instances ip addresses
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "sudo yum -y update"
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "sudo yum -y install gcc"
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "sudo yum -y install make gcc gcc-c++ kernel-devel"
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "sudo yum -y install golang"

# Install redis
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "wget http://download.redis.io/redis-stable.tar.gz"
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "tar -xzvf redis-stable.tar.gz"
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "cd ~/redis-stable && make distclean"
pssh -l ec2-user -h ips -x "-i cmpe281-project.pem" "cd ~/redis-stable && sudo make && sudo make install"
```

* Redis configuration
```
protected-mode no
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
appendonly yes
```

* Start Redis
```
./redis-server redis.conf
```

* Config Cluster Mode
The command used here is create, since we want to create a new cluster. The option --replicas 1 means that we want a slave for every master created. The other arguments are the list of addresses of the instances I want to use to create the new cluster.
```
gem install redis
./redis-trib.rb create --replicas 1 18.144.40.71:6379 52.53.251.248:6379 52.53.182.123:6379 18.144.73.22:6379 54.183.35.190:6379 54.183.41.186:6379
```

### 6. Kong Configuration (04/26/2018)
---
In this project, we use [Kong](https://getkong.org/about/) as our API gateway to route our API to different API servers.
We adopt Kong + PostgresSQL mode. PostgresSQL is deployed from DockerCloud, and Kong is deployed on AWS EC2 instance. In
Kong, we need to change few configs to make it talk to backend DB.
```
On 18.144.40.71
wget https://bintray.com/kong/kong-community-edition-aws/download_file?file_path=dists/kong-community-edition-0.13.1.aws.rpm
sudo yum install epel-release
sudo yum install kong-community-edition-0.13.1.aws.rpm --nogpgcheck

admin_listen = 0.0.0.0:8001, 0.0.0.0:8444 ssl  # Open to all foreign connections
change prefix to /home/ec2-user/kong
pg_host = <ec2 public IP>             # The PostgreSQL host to connect to.
pg_port = 5432                        # The port to connect to.
pg_user = kong                        # The username to authenticate if required.
pg_password = xxx                     # The password to authenticate if required.
pg_database = kong                    # The database name to connect to.
``` 

For PostgresSQL, we need to setup Kong user and database
```
On instance 54.183.192.182
psql -U postgres
CREATE USER kong; CREATE DATABASE kong OWNER kong;
```

```
kong migrations up [-c /path/to/kong.conf]
kong start [-c /path/to/kong.conf]
```

* Add Order API Routing
```
curl -i -X POST \
    --url http://18.144.40.71:8001/apis/ \
    --data 'name=order-api' \
    --data 'uris=/order' \
    --data 'strip_uri=true' \
    --data 'upstream_url=http://52.53.251.248:8080'
```

* Add User API Routing
```
curl -i -X POST \
    --url http://18.144.40.71:8001/apis/ \
    --data 'name=user-api' \
    --data 'uris=/user' \
    --data 'strip_uri=true' \
    --data 'upstream_url=http://18.188.63.102:8080'
```

* Add Inventory API Routing
```
curl -i -X POST \
    --url http://18.144.40.71:8001/apis/ \
    --data 'name=inventory-api' \
    --data 'uris=/inventory' \
    --data 'strip_uri=true' \
    --data 'upstream_url=http://{54.219.141.59}:8080'
```

* Add Review API Routing
```
curl -i -X POST \
    --url http://18.144.40.71:8001/apis/ \
    --data 'name=review-api' \
    --data 'uris=/reviews' \
    --data 'strip_uri=true' \
    --data 'upstream_url=http://54.218.74.6:8080'
```

### 7. Deploy API Server (04/20/2018)
Deploy GO API Server on each EC2 instance
Order System
```
On 52.53.251.248:8080
git clone https://github.com/nguyensjsu/team281-old-drivers.git
./dep.sh // import dependency
go install order
./bin/order <config_file> // config file includes redis masters' addresses: <ip>:<port>
```

User Management system
```
On 18.188.63.102:8080
git clone https://github.com/nguyensjsu/team281-old-drivers.git
./dep.sh // import dependency
go install user_management
./bin/user_management <config_file> // config file includes redis masters' addresses: <ip>:<port>
```

Inventory system
```
54.219.141.59:8080
git clone https://github.com/nguyensjsu/team281-old-drivers.git
./dep.sh // import dependency
go install inventory
./bin/inventory <config_file> // config file includes redis masters' addresses: <ip>:<port>
```

Review system
```
54.218.74.6:8080
git clone https://github.com/nguyensjsu/team281-old-drivers.git
./dep.sh // import dependency
go install review
./bin/review <config_file> // config file includes redis masters' addresses: <ip>:<port>

```

### Appendix
---
[Redis](https://redis.io/)
[Redis SDK](https://github.com/go-redis/redis)
[Kong](https://getkong.org/about/)
[Heroku](https://dashboard.heroku.com/)
