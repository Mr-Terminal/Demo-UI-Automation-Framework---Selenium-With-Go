
# Demo UI Automation framework - Selenium with Golang

This is a Demo Browser UI Automation framework written in Golang and used Selenium as Automation library and Go testing (unit) as the test framework.

The Folder Structure is something like below ->




## Folder Structure

![App Screenshot](https://i.postimg.cc/sX1VTGyy/image.png)

The Tests folder contain the test suites and grouped test cases respectively in a parent child manner.
PageObjects contain page object related methods and Utilty contains generic utility methods and scratch level browser UI interaction methods with Browserdriver.

TestMain_test.go is the entrypoint of the execution of the Automation. and common.go file contains generic methods thats gonna used only once while the whole module Automation Run.

go.mod is necessary for defining all dependencies and Module related data.




## Installation

Install necessary dependencies with go packages.
* Download GoLang
* Create Project Directory and move
* Create a module 

    ```bash
        go mod init
    ```
* Install Selenium & testify package

    ```bash
        go get -u github.com/tebeka/selenium
        go get github.com/stretchr/testify
    ```
    
## Running Tests

To run tests, run the following command

```bash
  go test -v
```
To run all the test files present in directory and subdirectories in the Module

```bash
  go test -v ./...
```
To run a particular package present in my Module

```bash
  go get -v ./Tests...
```
  


## Limitation

Currently in this Test Automation framework the execution order of the test is deprecated. Please find the related issue in issue section.

![App Screenshot](https://i.postimg.cc/yxMc1zng/image.png)

The Overall Flow should be like wise:
   * TestMain_test.go ---> setup()
   * m.Run()
   * TestMain_test.go ----> Teardown()

While calling m.Run() it can't able to detect the test files present inside Tests directory.   
## Feedback

If you have any feedback, please reach out to me at garai.sayantan10@gmail.com


