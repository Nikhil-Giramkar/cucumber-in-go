Feature: Basic Calculator
In this feature file we covered
multiple scenarios to test the basic Calculator

  @goimpl
  Scenario: Addition of 2 natural numbers
    When I click on digit 21
    And I click on Add button
    And i click on second digit 7
    Then calculation results are
      | Property | Value           |
      | Answer   |              28 |
      | Message  | Positive Number |

  @goimpl
  Scenario: Subtract of 2 natural numbers
    When I click on digit 12
    And I click on Subtract button
    And i click on second digit 3
    Then calculation results are
      | Property | Value           |
      | Answer   |               9 |
      | Message  | Positive Number |

  @goimpl
  Scenario: Multiplication of 2 natural numbers
    When I click on digit 11
    And I click on Multiply button
    And i click on second digit 11
    Then calculation results are
      | Property | Value           |
      | Answer   |             121 |
      | Message  | Positive Number |

  @goimpl
  Scenario: Division of 2 natural numbers
    When I click on digit 45
    And I click on Divide button
    And i click on second digit 15
    Then calculation results are
      | Property | Value           |
      | Answer   |               3 |
      | Message  | Positive Number |

  @goimpl
  Scenario: Invalid operation
    When I click on digit 45
    And I click on Exponent button
    And i click on second digit 15
    Then calculation results are
      | Property | Value             |
      | Answer   |                 0 |
      | Message  | Invalid Operation |
