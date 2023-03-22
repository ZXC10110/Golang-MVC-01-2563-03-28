package Controllers

import (
	"net/http"

	"github.com/zxc10110/mvc_13/Models"

	"github.com/gin-gonic/gin"
)

//GetAllHospitals ... Get all hospitals
func GetHospitals(c *gin.Context) {
	data, err := Models.GetAllHospitals()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

//GetAllPatients ... Get all patients
func GetPatients(c *gin.Context) {
	data, err := Models.GetAllPatient()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

//GetAllPatients ... Get all patients
func GetPatientCovidStatus(c *gin.Context) {
	data, err := Models.GetAllPatientCovidStatus()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

//GetAllPatientsInfo ... Get all patients
func GetCovidPatientInfo(c *gin.Context) {
	data, err := Models.GetCovidPatientInfo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

//GetAllPatientsInfo ... Get all patients
func GetCountPatient(c *gin.Context) {
	data, err := Models.GetCountPatientPerHospital()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

//GetAllPatientsInfo ... Get all patients
func GetOrderByCountPatient(c *gin.Context) {
	data, err := Models.GetOrderCountPatient()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

/*
//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
*/
