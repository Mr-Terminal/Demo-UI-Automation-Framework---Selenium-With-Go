package main

import (
	// "fmt"
	"github.com/tebeka/selenium"
	// "github.com/tebeka/selenium/chrome"
	"log"
	// "testing"
	"time"
)

var (
	BaseURL      = "https://opensource-demo.orangehrmlive.com/"
	LoginUserName     = "Admin"
	Password     = "admin123"
	Web_LoginBtn = `//button[text()=" Login "]`
	Web_TopNavProfileDropdown = `div.oxd-topbar-header-userarea > ul > li > span > i`
	Web_LogoutBtn = `//a[text()="Logout"]`
)

func CreateChromeInstance(){
	// initialize a Chrome browser instance on port 4444
	_, err := selenium.NewChromeDriverService("C:/Users/Sayantan/Downloads/chromedriver-win64/chromedriver-win64/chromedriver.exe", 4444)

	if err != nil {
		log.Fatal("Create Browser Instance Error:", err)
	}
}

func CreateRemoteClientWithCaps(caps selenium.Capabilities) selenium.WebDriver {
	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Create Remote Client Error:", err)
	}

	return driver
}

func LaunchURL(driver selenium.WebDriver) error {
	// navigate to the page
	err := driver.Get(BaseURL)
	driver.MaximizeWindow("")
	time.Sleep(5 * time.Second)
	return err
}

func LogIn(driver selenium.WebDriver) error {

	username, err := driver.FindElement(selenium.ByName, "username")
	if err != nil {
		return err
	}
	password, err := driver.FindElement(selenium.ByName, "password")
	if err != nil {
		return err
	}
	username.SendKeys(LoginUserName)
	password.SendKeys(Password)

	LoginBtn, err := driver.FindElement(selenium.ByXPATH, Web_LoginBtn)
	if err != nil {
		return err
	}

	LoginBtn.Click()
	time.Sleep(5 * time.Second)
	return nil
}

func LogOut(driver selenium.WebDriver) error{

	TopNavProfileDropdown, err := driver.FindElement(selenium.ByCSSSelector, Web_TopNavProfileDropdown)
	if err != nil {
		return err
	}
	TopNavProfileDropdown.Click()
	time.Sleep(2 * time.Second)

	LogoutBtn, err := driver.FindElement(selenium.ByXPATH, Web_LogoutBtn)
	if err != nil {
		return err
	}
	LogoutBtn.Click()
	time.Sleep(2 * time.Second)
	return nil
}
