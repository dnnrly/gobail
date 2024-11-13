Feature: Error cases

    Scenario Outline: Exits with error
        When the app runs with parameters "<parameter>"
        Then the app exits with an error
        And the app output does not contain "panic:"
        And the app output contains "exited with an error: there was an error"

        Examples:
            | parameter      |
            | exit-no-return |
            | exit-return    |
            | exit-return2   |


    Scenario Outline: Exits without adding error
        When the app runs with parameters "<parameter>"
        Then the app exits with an error
        And the app output does not contain "panic:"
        And the app output contains "exited with an error"
        And the app output does not contain "there was an error"

        Examples:
            | parameter                    |
            | exit-no-return-without-error |
            | exit-return-without-error    |
            | exit-return2-without-error   |

    Scenario Outline: Panics with error
        When the app runs with parameters "<parameter>"
        Then the app exits with an error
        And the app output contains "panic:"
        And the app output contains "exited with an error: there was an error"

        Examples:
            | parameter      |
            | panic-no-return |
            | panic-return    |
            | panic-return2   |


    Scenario Outline: Panics without adding error
        When the app runs with parameters "<parameter>"
        Then the app exits with an error
        And the app output contains "panic:"
        And the app output contains "exited with an error"
        And the app output does not contain "there was an error"

        Examples:
            | parameter                    |
            | panic-no-return-without-error |
            | panic-return-without-error    |
            | panic-return2-without-error   |
