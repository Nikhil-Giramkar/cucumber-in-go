Feature: Basic Calculator
In this feature file we covered
multiple scenarios to test the basic Calculator

  @goimpl
  Scenario: Addition of 2 natural numbers
    When I click on digit 12
    And I click on Add button
    And i click on second digit 7
    Then calculation results are
      | Property | Value           |
      | Answer   |              19 |
      | Message  | Positive Number |
