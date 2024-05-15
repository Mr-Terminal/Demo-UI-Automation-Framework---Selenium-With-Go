package utils

import (
	// "fmt"
	"log"
	"time"

	// "TestModule/config"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// It will set the driver instance for this package
func SetWebDriver(driverInstance *selenium.WebDriver){
	driver = *driverInstance
}

var (
	// driver                    = config.GetWebDriver()
	Web_SearchBarInputSideNav = `div > div > input`
	Web_SearchedFirstItem     = `//div[@class="oxd-sidepanel-body"]/ul/li/a/span`
)

func Search_and_Navigate_to_Other_Page_From_Side_Nav(SearchedText string) {
	elem, err := driver.FindElement(selenium.ByCSSSelector, Web_SearchBarInputSideNav)
	if err != nil {
		log.Fatal("No Such Element Found: ", err)
	} 
	
	elem.SendKeys(SearchedText)
	elem, err = driver.FindElement(selenium.ByXPATH, Web_SearchedFirstItem)
	if err != nil {
		log.Fatal("No Such Element Found: ", err)
	}
	elem.Click()
	time.Sleep(2 * time.Second)
}