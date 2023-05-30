# Go hands-on exercises

Repo created to practice Go-lang from the Udemy course
"Learn How To Code: Google's Go (golang) Programming Language"
from Todd McCleod.

<details>
<summary>Hands-on exercise #9</summary>

  - Create the following variables with the following scopes:
    - Package level
      - Create outside of `func main`
      - Use the
        - `var` keyword
        - `const` keyword
    - Block level
      - Inside `func main`
      - Use the short declaration operator
  - Use the variable in `func main`


</details>
<details>
<summary>Hands-on exercise #10</summary>

  - Use the terminal to make a Go workspace
    - `mkdir <name>`
    - `cd <name>`
    - `go mod init <some_name>`
  - Write a **hello world** program
    - `vim main.go`
    - Write Go code
  - Run your program
    - `go run main.go`

</details>
<details>
<summary>Hands-on exercise #11</summary>

  - Using the code you wrote in the previous hands-on exercise:
    - Look at the contents in the folder of your module
      - `ls -la`
    - Build your program
      - Any of these:
        - `go build main.go`
        - `go build .`
        - `go build ./...`
    - Run your executable:
      - `./<name of the executable>`

</details>
