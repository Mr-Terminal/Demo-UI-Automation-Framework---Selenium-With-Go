package testadmin

import (
	// "flag"
	"fmt"
	"testing"
	"time"

	"TestModule/PageObjects/admin"
	"TestModule/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// Suite Variables
var (
	// driver = config.GetWebDriver()
	SearchedText = "Admin"
	UserName = "Deepak Jacob"
	EmployeeName = "Ravi B"

)

// It will set the driver instance for this package
func SetWebDriver(driverInstance selenium.WebDriver){
	driver = driverInstance
}

// define suite with type struct and suite scoped variables
type AdminUserManagementSuite struct{
	suite.Suite
	// driver selenium.WebDriver 
}

// This will run before all the test cases in the suite are run
func (suite *AdminUserManagementSuite) SetupSuite() {
	fmt.Println("\t\t====== Suite Setup Called =======")
	utils.Search_and_Navigate_to_Other_Page_From_Side_Nav(SearchedText)
}

// This will run after all the test cases in the suite are completed
func (suite *AdminUserManagementSuite) TearDownSuite() {
	fmt.Println("\t\t====== Suite Teardown Called =======")
}

// This will run before each test cases
func (suite *AdminUserManagementSuite) SetupTest() {
	fmt.Println("\n\t\t\t====== Test Setup Called ======")
}

// This will run after each the test cases 
func (suite *AdminUserManagementSuite) TearDownTest(){
	fmt.Println("\n\t\t\t====== Test TearDown Called ======")
}


// This will run right before the test starts only if we invoke
// **** This function is not part of the standard testify package
// **** This function need to called manually in Test*** function if needed
func (suite *AdminUserManagementSuite) BeforeTest() {
	fmt.Println("===== BeforeTest Called =====")
}

// This will run right after the test finishes
// **** This function is not part of the standard testify package
// **** This function need to called manually in Test*** function if needed
func (suite *AdminUserManagementSuite) AfterTest() {
	fmt.Println("===== AfterTest Called =====")
}

// Tests related to User management in Admin Page
func (suite *AdminUserManagementSuite) Test1_Verify_that_Add_User_Btn_present_or_not_in_the_user_management_page(){

	var flag bool
	/*This test verifies if Add user Btn present in Admin User Management Page*/
	_, err := driver.FindElement(selenium.ByXPATH, admin.Web_AddUserBtn)
	if err != nil {
		flag = false
	} else {
		flag = true	
	}
	assert.True(suite.T(), flag, "Add User Button should be present")
}

func (suite *AdminUserManagementSuite) Test2_Verify_Search_functionality_in_user_management_page(){
	ele, err := driver.FindElement(selenium.ByCSSSelector, admin.Web_UserName_Input_Field)
	if err != nil {
		assert.Fail(suite.T(), "UserName Input Field is not present")
	}
	ele.SendKeys(UserName)

	// time.Sleep(3 * time.Second)

	// ele, err = driver.FindElement(selenium.ByCSSSelector, admin.Web_UserRole_Dropdown_Btn)
	// if err != nil {
	// 	assert.Fail(suite.T(), "UserRole Dropdown Button is not present")
	// }
	// ele.Click()

	time.Sleep(3 * time.Second)

	ele, err = driver.FindElement(selenium.ByCSSSelector, "body")
	if err!= nil {
		panic(err)
	}

	// Send the arrow down key
	ele.SendKeys(string(selenium.DownArrowKey))

	ele.SendKeys(string(selenium.EnterKey))

	time.Sleep(3 * time.Second)

	ele, err = driver.FindElement(selenium.ByXPATH, admin.Web_EmployeeName_Input_Field)
	if err != nil {
		assert.Fail(suite.T(), "Employee Input Field is not present")
	}
	ele.SendKeys(EmployeeName)

	time.Sleep(3 * time.Second)

	ele, err = driver.FindElement(selenium.ByCSSSelector, admin.Web_Status_Dropdown_Btn)
	if err != nil {
		assert.Fail(suite.T(), "Status Dropdown Button is not present")
	}
	ele.Click()

	time.Sleep(3 * time.Second)

	ele, err = driver.FindElement(selenium.ByCSSSelector, "body")
	if err!= nil {
		panic(err)
	}

	// Send the arrow down key
	ele.SendKeys(string(selenium.DownArrowKey))

	ele.SendKeys(string(selenium.EnterKey))

	time.Sleep(3 * time.Second)

	ele,err = driver.FindElement(selenium.ByXPATH, admin.Web_SearchBtn)
	if err != nil {
		assert.Fail(suite.T(), "Search Button is not present")
	}
	ele.Click()

	time.Sleep(3 * time.Second)

}

// This will initiate the Suite Run
func TestRunSuite(t *testing.T){
	fmt.Println("\t===== Suite Run Starts =====")
	suite.Run(t, new(AdminUserManagementSuite))
}
