package Controllers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zxc10110/mvc_13/View"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("getAllHospitals", GetHospitals)
	r.GET("getAllPatients", GetPatients)
	r.GET("getAllPatientsCovidStatus", GetPatientCovidStatus)
	r.GET("getCovidPatientInfo", GetCovidPatientInfo)
	r.GET("getCountPatient", GetCountPatient)
	r.GET("getOrderByCountPatient", GetOrderByCountPatient)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
