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
<details>
<summary>Hands-on exercise #12</summary>

  - Using the code you wrote in the previous hands-on exercise:
    - Build your program for Windows
      - `GOOS=windows go build -o main-windows.exe main.go`
    - Build your program for Mac
      - `GOOS=darwin go build -o main-mac main.go`
    - Build your program for Linux
      - `GOOS=linux go build -o main-linux main.go`

</details>
<details>
<summary>Hands-on exercise #13</summary>

  - Using the code you wrote in the previous hands-on exercise:
    - Make sure your `$GOPATH/bin` is part of your `$PATH` envariable
    - Look in the `$GOPATH/bin`
      - Launch another terminal
      - See the `GOPATH` environment variable with either of these:
        - `go env`
        - `go env GOPATH`
      - Navigate to the `$GOPATH/bin` folder
      - `ls -la`
    - `go install` your program (on the first/other terminal)
      - Look at the executable `$GOPATH/bin`
    - Run the executable in the `$GOPATH/bin`
    - Remove the executable in the `$GOPATH/bin`
      - If you accidentally delet everything, you will need to reinstall your tooling in VScode
      - If you messed it all up, reinstall go.

</details>
<details>
<summary>Hands-on exercise #14</summary>

  - Using the code from the previous hands-on exercise:
    - Use a function from the package found at **github.com/GoesToEleven/puppy**
      - `go get github.com/GoesToEleven/puppy`
    - Inspect your `go.mod` file
    - Run `go mod tidy`
    - What does `go mod tidy` do?
      - https://go.dev/ref/mod#go-mod-tidy

</details>
