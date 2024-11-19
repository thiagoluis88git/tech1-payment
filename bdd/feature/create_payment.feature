Feature: Pay order
  In order to get the order done
  As a customer of the fastfoot restaurant
  I need to be able to pay the amount bill

  Scenario: then user try to pay the order, success should be displayed
    When I send "POST" request to "/books" with payload:
      """
      {
          "id": 1,
          "title": "Dune",
          "author": "Frank Herbert"
      }   
      """
    Then the response code should be 201
    And the response payload should match json:
      """
      [
          {
              "id": 1,
              "title": "Dune",
              "author": "Frank Herbert"
          }
      ]   
      """