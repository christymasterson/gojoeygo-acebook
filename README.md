# gojoeygo-acebook

## Summary:
Welcome to the repository for our first engineering project at Makers Academy. In this project we aim to create a social networking web app 'acebook', written in Go.

We tried to keep our application as simple as possible. You can login, logout and signup as a new user. You can create posts and see posts in reverse chronological order with timestamp and user name who created the post. There is a like feature added also where user can only like post once.

## The Team:
Our team 'Go Joey Go' is made up of four apprentices, 8 weeks into our 12 week bootcamp. For all of us, this will be our first time programming in Go language.

## Our Objectives:
Our primary objective is to TDD a functional web app in Go.
In addition to this we aim to document the process of learning a new language and develop our teamwork skills.

## Technologies Used:
Primary programming language: Go
Front end languages: HTML, CSS
Text editors: Atom / VS Code
Web framework: Gin
Testing framework: Go’s inbuilt testing functionality
Database: PostgreSQL, PQ, TablePlus, SQL
Version control: GitHub
Hosting platform: Heroku

## Database Setup:
When working with this repo for the first time you must set up the database on your machine. We assume here that Postgres is installed.
To set up the database:
1. From the command line run psql.
2. Execute the commands stored in db/migrations.

## How to install and run the app:
   1. Make sure you have installed 'golang' on you computer.
  - [Go Website](http://www.golang.org/dl/)
   2. clone the project repository in the default golang path which is $HOME/go/src
   3. Run the 'dep ensure' command from inside the project folder to make sure all the dependencies are installed.
   4. Set up the databased as mentioned above.
   5. Run the app by calling ./app from the command-line.

## How to run tests:
The following commands can be run from command-line:
- go test

You can output print additional information about test function using verbose -v command:
- go test -v

Command fo checking test coverage:
- go test -cover

Command to generate detailed test coverage report:
- go test -coverprofile=coverage.out

Command for generating the coverage profile in HTML file:
- go tool cover -html=coverage.out

## Process followed to learn Go language:
1. Learned golang by translating the projects we did in Ruby language.
2. Used the resources like 'Codecademy' and 'A Tour of Go'.
3. Tried to learn by TDD approach.
4. Learned by trial and error.
5. Followed debugging approach.
6. Reaching out for support.
7. Kept the thing as simple as we could understand.
8. We used pairing technique to learn and implement new feature.
9. Learned by sharing knowledge and resources.
