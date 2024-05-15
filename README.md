
# Demo UI Automation framework - Selenium with Golang

This is a Demo Browser UI Automation framework written in Golang and used Selenium as Automation library and Go testing (unit) as the test framework.

The Folder Structure is something like below ->




## Folder Structure

![App Screenshot](https://i.postimg.cc/sX1VTGyy/image.png)

The Tests folder contain the test suites and grouped test cases respectively in a parent child manner.
PageObjects contain page object related methods and Utilty contains generic utility methods and scratch level browser UI interaction methods with Browserdriver.

TestMain_test.go is the entrypoint of the execution of the Automation. and common.go file contains generic methods thats gonna used only once while the whole module Automation Run.

go.mod is necessary for defining all dependencies and Module related data.



