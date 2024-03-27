# Cucumber in GoLang

Integration testing ensures that different components or modules of your application work together correctly.<br>

BDD testing is an Agile approach to software testing where testers write test cases in simple language that even people without technical expertise can understand <br>

Cucumber is a Behavior Driven Development (BDD) tool used to develop test cases to verify if a particular feature in our application functions correctly <br>

All scenarios to be covered and background setup for the tests ae written in Gherkin format <br>

# About the project

This project demonstrates how Cucumber is used to test features in GoLang. <br>
To use Cucumber framework I made use of 2 important packages
    <li>goDog - https://pkg.go.dev/github.com/cucumber/godog</li>
    <li>assistDog - https://pkg.go.dev/github.com/rdumont/assistdog</li>

Here, we tried to validate different scenarios of a basic calculator <br>
Actual Values are returned from Backend as per logic implemented by Dev team<br>
Those are validated against the expected values written in .feature file by QA Team<br>

# To run the project

Open the folder in integrated terminal in VS-Code
Run the following commands
```
    go mod tidy
    go test -v -run ^TestData$
```
This will give results as below
![Screenshot (311)](https://github.com/Nikhil-Giramkar/cucumber-in-go/assets/58767494/423b376d-989b-4900-a117-c0d4c32d9ebb)
