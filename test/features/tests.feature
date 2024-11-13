Feature: Error cases

    Scenario: Default case does not error (validate testapp)
        When the app runs without parameters
        Then the app exits without error
        And the app output has 1 line

    Scenario Outline: Just exits
        When the app runs with parameters "<parameter>"
        Then the app exits with an error
        And the app output has 1 line

        Examples:
            | parameter             |
            | exit-no-return-no-msg |
            | exit-return-no-msg    |
            | exit-return2-no-msg   |

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

    Scenario Outline: Just panics
        When the app runs with parameters "<parameter>"
        Then the app exits with an error
        And the app output contains "panic:"
        And the app output has 2 lines

        Examples:
            | parameter              |
            | panic-no-return-no-msg |
            | panic-return-no-msg    |
            | panic-return2-no-msg   |

    Scenario Outline: Panics with error
        When the app runs with parameters "<parameter>"
        Then the app exits with an error
        And the app output contains "panic:"
        And the app output contains "exited with an error: there was an error"

        Examples:
            | parameter       |
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
            | parameter                     |
            | panic-no-return-without-error |
            | panic-return-without-error    |
            | panic-return2-without-error   |
