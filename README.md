# Golang-simple-rest-api


This is an example of implementation of Clean Architecture in Go (Golang) projects.


From Uncle Bob’s Architecture we can divide our code in 4 layers :


Entities: encapsulate enterprise wide business rules. An entity in Go is a set of data structures and functions.

Use Cases: the software in this layer contains application specific business rules. It encapsulates and implements all of the use cases of the system.

Controller: the software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web

Framework & Driver: this layer is generally composed of frameworks and tools such as the Database, the Web Framework, etc.


I used 4 layer:

Models <=> Entities

Services <=> Use Cases :  -- I have a base sevice serve CRUD  all dto

Controler :

    + -- Controller have 2 sevice : BaseService and DTOService 
    + -- On Controller works : Validate value form client => Check permit access => Use function on service ,...
    
Repository <=> Framework & Driver : -- I use GORM v1 framework


The diagram:


<p><a target="_blank" rel="noopener noreferrer" href="https://github.com/nvt206/Golang-simple-rest-api/blob/master/clean-arch.png"><img src="https://github.com/nvt206/Golang-simple-rest-api/blob/master/clean-arch.png" alt="golang clean architecture" style="max-width:100%;"></a></p>

The original explanation about this project's structure can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047.

More at: https://github.com/bxcodec/go-clean-arch

