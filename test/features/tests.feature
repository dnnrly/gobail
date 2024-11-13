Feature: Error cases

    Scenario: Exits gobail.Run
        When the app runs with parameters "exit-no-return"
        Then the app exits with an error
        And the app output contains "exited with an error: there was an error"

    Scenario: Exits gobail.Run without adding error
        When the app runs with parameters "exit-no-return-without-error"
        Then the app exits with an error
        And the app output contains "exited with an error"
        And the app output does not contain "there was an error"

    Scenario: Exits gobail.Return
        When the app runs with parameters "exit-return"
        Then the app exits with an error
        And the app output contains "exited with an error: there was an error"

    Scenario: Exits gobail.Return without adding error
        When the app runs with parameters "exit-return-without-error"
        Then the app exits with an error
        And the app output contains "exited with an error"
        And the app output does not contain "there was an error"
        