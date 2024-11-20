Feature: Pay order
  In order to get the order done
  As a customer of the fastfoot restaurant
  I need to be able to pay the amount bill

  Scenario: then user try to pay the order, success should be displayed
    When I send "POST" request to "/payment" with payload:
      """
      {
          "totalPrice": 123.32,
          "paymentType": "CREDIT"
      }   
      """
    Then the response code should be 200
    And the response payload should match json:
      """
        {
            "paymentId": "rwer342534sdf",
            "paymentGatewayId": "234trr00"
        } 
      """