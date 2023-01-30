# Car Rental Go 




To Run this project locally without Docker 

1. clone repository 
2. Run `go mod tidy`  to install go dependancies
3. Setup postgres locally, Create user, setup password, create database, grant database access to user
4. Update .env file with username, password and DB_name and change the DB_HOST = localhost
5. Run `go run cmd/main.go`  

With Docker : 
Run `docker compose up` 
to stop the db containers run `docker compose down`

Use Postmen for testing API's



1. SIGN UP NEW USER

- (POST request) url : `http://localhost:3000/user/signup` 

- Add below JSON in the the request body 

`
{
    "first_name": "userF",
    "last_name" : "userL",
    "email"     : "user@gmail.com",
    "phone"     : "1234567890",
    "password"  : "password",
    "user_type" : "ADMIN"       

}

`
- you can change user_type to `USER` 



2. LOGIN USER

(POST request ) url : `http://localhost:3000/user/login`

to login user, add following json to request body 
` 
{
    "email"     : "user@gmail.com",
    "password"  : "password "
 }

 `
 you will get user info as Response 
 copy jwt-token from field "token" 



3. Add a new Car to database
 
- url :`http://localhost:3000/cars/AddCar`

- Add below JSON in the request body

`
{
    "carModel": "fc-001",
    "date_of_manufacture": "1-04-2021",
    "last_service_date" :"6-2-2022",
    "last_used_date" : "4-1-2023",
    "status" : "Rented"

}
`

!!! DONT FORGET TO ADD JWT-TOKEN BEFORE SENDING REQUEST !!!

- Add jwt-token copied from user info response to the header of the request
`
key = token 
value = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJVc2VyX3R5cGUiOiIiLCJleHAiOjE2NzU0NDI5MTd9.E8vq-bnWrUjiOoTqgbYU9Ma1AEM_oQh9_oLpZtlfk1A`

`



4. GET ALL CAR

- (GET request )url : `http://localhost:3000/cars`

       OR
       
- Graphql querry := 
`
query GetAllCars
    { 
        cars{ 
            id 
            car_model
            date_of_manufacture
            last_service_date
            last_used_date
            status                
            } 
    }
`
- !!! DONT FORGET TO ADD JWT-TOKEN BEFORE SENDING REQUEST !!!
- You will get list of all cars in the database

- Copy the id of any car you want to rent/book



5. Book Car 

-  (GET request) url : `http://localhost:3000/cars/bookCar/:carID`

- Add the carID in the request
 example : `http://localhost:3000/cars/bookCar/4ef1af7e-3076-4c29-9358-4365cd42f0d8`
 
 

!!! DONT FORGET TO ADD JWT-TOKEN BEFORE SENDING REQUEST !!!


6. Return Car 

- (GET request) url : `http://localhost:8080/cars/returnCar/:carID`

- Add the carID in the request
 example : `http://localhost:3000/cars/returnCar/4ef1af7e-3076-4c29-9358-4365cd42f0d8`


!!! DONT FORGET TO ADD JWT-TOKEN BEFORE SENDING REQUEST !!!



7. LOGOUT user

- (GET request) url : `http://localhost:8080/users/logout`

!!! DONT FORGET TO ADD JWT-TOKEN BEFORE SENDING REQUEST !!!
