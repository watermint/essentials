# Essential libraries for Go

# Goals

* All native Go library.
* Minimize `if err != nil` idiom. Compare to `nil` sometime cause unintended bugs.
* Minimize use of external libraries.
* Separate interface from implementation.
* Ease of development.

# Non-goals

* Performance (this is often trade-off to ease-of-development)
