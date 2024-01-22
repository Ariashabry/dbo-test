# TechnicalTest Golang Depoguna Bangunan Online

Name : Aria shabry (aria.shabry@gmail.com)



**Setting Environtment**

1. Update your environtment at env.yaml file like DB_Host, DB_PORT, DB_NAME, and etc.
2. If your environtment error, please check env.yaml file and ConnectionString at ./helpers/env/env.go
3. Database Name = dbo_test
4. Import Database using dbo_test.sql file
5. Application address : http://localhost:5000/

**Answer**
Project ini hanya simulasi dan cuma memuat MVP



1. Untuk ERD
./Documentation/erd.jpg
![ERD](https://github.com/Ariashabry/dbo-test/blob/main/documentation/erd.jpg?raw=true)

2. File sql (I'm using MySQL)
./Documentation/dbo_test.sql
https://github.com/Ariashabry/dbo-test/blob/main/documentation/dbo_test.sql

3. Change your Environtment at env.yaml
./env.yaml
https://github.com/Ariashabry/dbo-test/blob/main/env.yaml

4. Untuk Dockerfile 
./DockerFile
https://github.com/Ariashabry/dbo-test/blob/main/Dockerfile


5. Untuk Postman Collection
./Documentation/Test.postman_collection.json
https://github.com/Ariashabry/dbo-test/blob/main/documentation/DBO%20Test.postman_collection.json

6. API Documentation
https://app.swaggerhub.com/templates/ARIASHABRY_1/DBO-Online/1.0.2#

6. Berikut Saya Tampilkan Screenshoot hasil running di local

- POST : http://localhost:5000/customer
![Insert](https://github.com/Ariashabry/dbo-test/blob/main/documentation/1.%20customer%20-%20create.png?raw=true)

- GET : http://localhost:5000/customer?length=10&start=1
![GetALL](https://github.com/Ariashabry/dbo-test/blob/main/documentation/2.%20customer%20-%20get%20All.png?raw=true)

- GET : http://localhost:5000/customer/7
![3](https://github.com/Ariashabry/dbo-test/blob/main/documentation/3.%20customer%20-%20get%20By%20id.png?raw=true)

- PUT : http://localhost:5000/customer
![4](https://github.com/Ariashabry/dbo-test/blob/main/documentation/4.%20customer%20-%20update.png?raw=true)

- DELETE : http://localhost:5000/customer/10
![5](https://github.com/Ariashabry/dbo-test/blob/main/documentation/5.%20customer%20-%20delete.png?raw=true)

- POST : http://localhost:5000/order
![6](https://github.com/Ariashabry/dbo-test/blob/main/documentation/6.%20Order%20-%20create.png?raw=true)

- GET : http://localhost:5000/order?length=10&start=1
![7](https://github.com/Ariashabry/dbo-test/blob/main/documentation/7.%20Order%20-%20get%20All.png?raw=true)

- GET : http://localhost:5000/order/6
![8](https://github.com/Ariashabry/dbo-test/blob/main/documentation/8.%20Order%20-%20get%20By%20Id.png?raw=true)

- PUT : http://localhost:5000/order
![9](https://github.com/Ariashabry/dbo-test/blob/main/documentation/9.%20Order%20-%20Update.png?raw=true)

- DELETE : http://localhost:5000/order/7
![10](https://github.com/Ariashabry/dbo-test/blob/main/documentation/10.%20Order%20-%20Delete.png?raw=true)

- POST : http://localhost:5000/login
![11](https://github.com/Ariashabry/dbo-test/blob/main/documentation/11.%20Login.png?raw=true)

- POST : http://localhost:5000/register
![12](https://github.com/Ariashabry/dbo-test/blob/main/documentation/12.%20Register.png?raw=true)


Jika ada pertanyaan, silahkan hubungi saya di : aria.shabry@gmail.com


