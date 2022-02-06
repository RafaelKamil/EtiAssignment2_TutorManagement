
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/RafaelKamil/EtiAssignment2_TutorManagement">
  </a>

<h3 align="center">ETI Assignment 2</h3>

  <p align="center">
    Package 3.3
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
This project is a component of the school's management system. My role is tutor management, which includes creating, viewing, updating, and deleting tutor accounts.
Info includes
Tutor ID, Name, Description, Tutor ID, Name, Description,

### Built With

* [golang](https://go.dev/)
* [JavaScript](https://www.javascript.com/)
* [HTML](https://en.wikipedia.org/wiki/HTML)
* [Bootstrap](https://getbootstrap.com/)
* [MySql](https://www.mysql.com/)


<!-- GETTING STARTED -->
## Design consideration the microservice

The Tutor Management system is divided into three parts: frontend, database, and backend. The Frontend is reliant on the Backend, and the Backend is reliant on the Frontend.


The Database serves as a kind of persistent storage for the data that the Backend requires.

The API on the backend supports the four fundamental HTTP methods: GET, POST, PUT, and DELETE. To make the backend as modular as feasible, this was done. Although the Tutor Management backend presently has no external dependencies, its simplicity will make future integration with other microservices easy.


The Frontend is where the user interface and client-side logic are kept. The Frontend must hold the majority of the sophisticated functionality, including sending GET and PUT requests to the Class Microservice, in order for the Backend to be loosely connected. One significant disadvantage is security, as storing much of your functionality on the client-side exposes it to potential attackers. This is something that subsequent designs can indeed enhance.



### Dabase


Mysql is the database I use. The reason for this is because mysql includes an optional component called the MySQL HTTP Plugin. This plugin provides direct access to MySQL through a REST over HTTP interface, removing the requirement for a middle-tier server or database-specific drivers.

### Gorm
To handle all Mysql queries, such as creating table ,column and POST GET PUT DELETE . I'm using the Gorm package so i don't need to write any querys

What is the purpose of Gorm? 

GORM provides CRUD operations and can also be used for initial migration and schema construction. GORM also excels in its extendability with native plugin support, reliance on testing, and one-to-one and one-to-many group of associations.

###  MUX
The term mux is an abbreviation for "HTTP request multiplexer." ServeMux, mux, mux, mux, mux, mux, mux, mux, mux, mux, Incoming requests are compared to a list of registered routes, and the route that matches the URL or other conditions is called.

What's the point of Gorilla mux?
The gorilla/mux package includes a request router and dispatcher for routing inbound requests to the appropriate handler. The term mux is an abbreviation for "HTTP request multiplexer."
### Validations


To validate user input for the creating new passengers, drivers, and trips. I'm use the Validator Package.

What exactly is the Validator Package?

Package validator uses tags to implement value validations for structs and individual fields. It also supports Cross-Field and Cross-Struct validation for nested structs and can dive into arrays and maps of any type.


### Microservice Design
![assignment 2 design drawio (1)](https://user-images.githubusercontent.com/74031156/152704678-46ea6790-1f3f-4290-b0c2-3388697a798b.png)



### Struct Design
![image](https://user-images.githubusercontent.com/74031156/152705339-77aff38b-2c43-4ad9-bb8a-0e5aa0950699.png)


Architecture diagram
### Installation
From github
1. Clone the repo git clone https://github.com/RafaelKamil/EtiAssignment2_TutorManagement To be Advise Install the following libraries for each microservice: Trips, Passenger, and Driver

2. Install the following packages 
    go get -u github.com/gorilla/mux
    
    go get -u gorm.io/gorm
    
    go get -u gorm.io/driver/mysql
    
    go get github.com/go-playground/validator/v10
    
    go get "github.com/go-sql-driver/mysql"
    
From Dockerhub
1. Clone the repo From my Github 

  Microservice Backend
     <li> docker pull rafaelkamil/assignment2_modulemanagementcontainer</li>
  
  Front end
    <li> docker pull rafaelkamil/assignment2_backend_tutormanagement</li>
 
 Database
    <li>docker pull rafaelkamil/assignment2_database_tutormanagement</li>


2. Install the following packages 
    go get -u github.com/gorilla/mux
    
    go get -u gorm.io/gorm
    
    go get -u gorm.io/driver/mysql
    
    go get github.com/go-playground/validator/v10
    
    go get "github.com/go-sql-driver/mysql"
    
 
3. Build and run docker images using docker-compose
    
    Assignment-2\DockerAssignment-2 docker-compose> run --build



<!-- USAGE EXAMPLES -->
## How to run the program

 Assignment-2\DockerAssignment-2 docker-compose> run --build


<!-- ROADMAP -->
## Roadmap

- [] Create Databas Schema and the account
- [] Create the microservices (Tutor Management)
- [] Create front end 

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Your Name - [@Email Addres ] - isyhak98@gmail.com


