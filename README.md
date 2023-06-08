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
<details>
<summary>Hands-on exercise #15</summary>

  - Using the code from the previous hands-on exercise:
    - Use a function from the package found at **github.com/GoesToEleven/puppy** but make your code depend on v1.2.0
      - `go get github.com/GoesToEleven/puppy@v1.2.0`
    - Inspect your `go.mod` file

</details>
<details>
<summary>Hands-on exercise #16</summary>

  - Create a github repo for your code
  - Push your code
    - `git add -A`
    - `git commit`
      - Add a comment significant to your commit
    - `git push`
  - Add a version tag to your code of `v1.0.0`
    - `git tag`
    - `git tag v1.0.0`
    - `git push origin --tags`
  - Add a `temp.txt` file to your code
  - Push your code again
  - Add a version tag to your code of `v1.1.0`
  - Look at your version in github
  - Optional: delete your repository

</details>
<details>
<summary>Hands-on exercise #17</summary>

  - At the terminal:
    - go to your go workspace where you wrote your code
    - remove the folder you created to write your go code
      - rm -rf <folder name>

</details>
<details>
<summary>Hands-on exercise #18</summary>

  - This hands-on exercise has a text file associated with it.
  - When I ran a SHA-256 checksum on this text file, I got this hash:
    - `7c6c8937b2a120af15849db05c9f46326761e0eec852a2e973b1e0b6acd59a01`
  - Download the text file associated with this hands-on exercise. Run a SHA-256 checksum on it. Do you get the same hash?
    - shasum -a 256 `/path/to/file`
      - `.` is the current directory
      - if you're in the directory with the file
        - `shasum -a 256 ./file`
  - change the file by one character, then run SHA-256 again
    - `9be13f9173f28ce3dd89c72aad7f5b0549a0641feb869509c7f96e8dc8b6ea8e`

</details>
<details>
<summary>Hands-on exercise #19</summary>

  - create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your program should do the following
    - create a random int between 0 and 250
    - store the value of that int in a variable with the identifier of x
      - `func Intn(n int) int`
        - `rand.Intn()`
    - print out the name and value of the variable
  - use an if statement to do the following
    - if the value is between 0 and 100
      - print between 0 and 100
    - if the value is between 101 and 200
      - print between 101 and 200
    - if the value is between 201 and 250
      - print between 201 and 250

</details>
<details>
<summary>Hands-on exercise #20</summary>

  - Modify the previous program to use one of these conditional logic statements
    - a switch statement
    - a select statement
  - Which of the above conditional logic statements did you choose and why?

</details>
<details>
<summary>Hands-on exercise #21</summary>

  - Modify the previous program to have your program use the init func to print
    - "This is where initialization for my program occurs"
      - read about the init function in effective go
    - What does niladic mean?

</details>