package main

import (
	// "fmt"
	// "TestModule/Tests/admin"
	// "TestModule/config"
	"fmt"
	"log"
	"testing"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

	"TestModule/utils"
	// "TestModule/Tests/testadmin"
	"TestModule/PageObjects/admin"

)

var driver selenium.WebDriver

func TestMain(m *testing.M) {

	CreateChromeInstance()

	// configure the browser options
	caps_val := selenium.Capabilities{}
	caps_val.AddChrome(chrome.Capabilities{Args: []string{
		// "--headless", // comment out this line for headed mode
	}})

	driver = CreateRemoteClientWithCaps(caps_val)
	
	// share the driver instance to tests, pageObjects and utility layer
	utils.SetWebDriver(&driver)
	admin.SetWebDriver(&driver)
    // testadmin.SetWebDriver(&driver)

	// visit the target page
	err := LaunchURL(driver)
	if err != nil {
		log.Fatal("Launch URL Error:", err)
	}

	//Login to application
	err = LogIn(driver)
	if err != nil {
		log.Fatal("Login Error:", err)
	}


	// Starts to run the test suites present in this module
	fmt.Print("Before m.Run() !\n")
	defer m.Run()
	fmt.Print("After m.Run() !\n")

	
	//Logout to application
	err = LogOut(driver)
	if err != nil {
		log.Fatal("Logout Error:", err)
	}

	defer driver.Close()
}
