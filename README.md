# Quiz Client-Side CLI

This repository contains the client-side code for a Quiz application.

## Instructions to Install

To install the Quiz client, follow these steps:

1. **Get Dependencies**
    ```bash
    go get .
    ```

2. **Build Application**
    ```bash
    go build
    ```
    
3. **Install Application**
    ```bash
    go install
    ```

## Instructions to Use

To use the Quiz client, you can run 3 different commands:

1. **Get Questions**
    ```bash
    quiz-client-cli get-questions
    ```

2. **Submit Answers**
    ```bash
    quiz-client-cli submit-answers -u {user} -a {csv list of all answers in sequentical order}
    ```
    Example:
    ```bash
    quiz-client-cli submit-answers -u john -a 3,1,1,0,2,1,0,2,3,1,3,0,0,2,1,2,3,3,2,0
    ```
    
3. **Get User Statistics**
    ```bash
    quiz-client-cli get-user-statistics -u {user}
    ```
    Example:
    ```bash
    quiz-client-cli get-user-statistics -u john
    ```

