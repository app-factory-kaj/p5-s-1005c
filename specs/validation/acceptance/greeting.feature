Feature: Greeting

  @story-1
  Rule: Calling GET /hello with a name returns a JSON greeting addressed to that name

    Scenario: A calling service greets "Ada"
      Given the greeter service is available
      When a calling service calls GET /hello with name "Ada"
      Then it receives a JSON greeting addressed to "Ada"

    Scenario: A calling service greets "Grace"
      Given the greeter service is available
      When a calling service calls GET /hello with name "Grace"
      Then it receives a JSON greeting addressed to "Grace"

  @story-2 @negative
  Rule: Calling GET /hello without a valid name is refused with a clear error

    Scenario: The name parameter is missing
      Given the greeter service is available
      When a calling service calls GET /hello without a name parameter
      Then it receives an error response and no greeting

    Scenario: The name parameter is empty
      Given the greeter service is available
      When a calling service calls GET /hello with an empty name parameter
      Then it receives an error response and no greeting
