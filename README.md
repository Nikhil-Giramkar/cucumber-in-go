# Cucumber in GoLang

Integration testing ensures that different components or modules of your application work together correctly.<br>

BDD testing is an Agile approach to software testing where testers write test cases in simple language that even people without technical expertise can understand <br>

Cucumber is a Behavior Driven Development (BDD) tool used to develop test cases to verify if a particular feature in our application functions correctly <br>

All scenarios to be covered and background setup for the tests ae written in Gherkin format <br>

# About the project

This project demonstrates how Cucumber is used to test features in GoLang. <br>
Here, we made use of 2 important packages 
<ul>
    <li>goDog - https://pkg.go.dev/github.com/cucumber/godog</li>
    <li>assistDog - https://pkg.go.dev/github.com/rdumont/assistdog<li>
</ul>

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
