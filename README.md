# Cinema ticket booking
Create a system that allows users to book a seat in a movie theatre - no authentication required.
This assignment is mainly to test the skills of the developers which would be useful for e-commerce websites or applications like shopping carts, booking systems and the like.

Display all the seats in the theater and allow users to book it by clicking it. Only one user should be allowed to reserved a specific seat.
If another user clicks a seat that was booked, he should get an error. You must handle the concurrency scenarios and avoid data inconsistency.

Seats can be unbooked by clicking the booked seat again.

## Prerequisites
    Need to have 'go' installed on your machine. 

## Run on your machine

First setup config.json file with all your credentials. Import mysql database 'cinema_tkt'.

Install dependencies packages by following commands

    go get github.com/acoshift/flash
    go get github.com/codegangsta/negroni
    go get github.com/go-sql-driver/mysql
    go get github.com/gorilla/mux
    go get github.com/gorilla/sessions
    go get github.com/jmoiron/sqlx



Set up variables GOROOT, GOPATH and PATH and build project by following command in terminal, this will create an executable build file. To see output run executable file

    go build

and then run executable file

then visit in browser

    http://localhost:7000/

## For test cases run this command

    go test -v
