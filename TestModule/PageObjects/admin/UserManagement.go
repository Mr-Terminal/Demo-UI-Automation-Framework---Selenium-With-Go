package admin

import (
	// "fmt"
	// "testing"
	// "fmt"
	"time"

	// "TestModule/utils"

	// "github.com/stretchr/testify/assert"
	"log"

	// "github.com/stretchr/testify/suite"
	"github.com/tebeka/selenium"
)	

var driver selenium.WebDriver

func SetWebDriver(driverInstance *selenium.WebDriver){
	driver = *driverInstance
}

var (
	Web_AddUserBtn = `//button[text()=" Add "]`
	Web_UserName_Input_Field = "div:nth-child(1) > div > div:nth-child(2) > input"
	Web_UserRole_Dropdown_Btn = "div:nth-child(2) > div > div:nth-child(2) > div div.oxd-select-text--after"
	Web_EmployeeName_Input_Field = `//input[@placeholder="Type for hints..."]`
	Web_Status_Dropdown_Btn = "div:nth-child(4) > div > div:nth-child(2) > div > div > div.oxd-select-text--after"
	Web_SearchBtn = `//button[text()=" Search "]`
)

func Click_On_Add_User_Btn(){
	ele, err := driver.FindElement(selenium.ByXPATH, Web_AddUserBtn)
	if err != nil {
		log.Fatal("No Such Element: ", err)
	}
	ele.Click()
	time.Sleep(5 * time.Second)
}

