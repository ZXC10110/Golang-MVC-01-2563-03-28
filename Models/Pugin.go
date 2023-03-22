package Models

import (
	"github.com/zxc10110/mvc_13/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllHospitals
func GetAllHospitals() (hospitals []Hospitals, err error) {
	if err = Config.DB.Find(&hospitals).Error; err != nil {
		return
	}
	return
}

//GetAllHospitals
func GetAllPatient() (patients []Patients, err error) {
	if err = Config.DB.Find(&patients).Error; err != nil {
		return
	}
	return
}

//GetAllHospitals
func GetAllPatientCovidStatus() (covidStatus []PatientCovidStatus, err error) {
	if err = Config.DB.Select("*").
		Table("Test.patient_covid_status").
		Find(&covidStatus).Error; err != nil {
		return
	}
	return
}

//GetPatientInfo
func GetCovidPatientInfo() (covidPatient []AllCovidPatientInfo, err error) {
	if err = Config.DB.Select("c.hnid, p.first_name, p.last_name, c.covid_status").
		Table("Test.patient_covid_status as c").
		Joins("left join patients as p on p.hnid = c.hnid").
		Find(&covidPatient).Error; err != nil {
		return
	}
	return
}

//CountPatient
func GetCountPatientPerHospital() (count []CountPatient, err error) {
	if err = Config.DB.Select("COUNT(p.hnid) as count, h.hid, h.title").
		Table("Test.hospitals as h").
		Joins("left join patients as p on p.hid = h.hid").
		Group("hid, title").
		Find(&count).Error; err != nil {
		return
	}
	return
}

//OrderCountPatient
func GetOrderCountPatient() (count []CountPatient, err error) {
	if err = Config.DB.Select("COUNT(p.hnid) as count, h.hid, h.title").
		Table("Test.hospitals as h").
		Joins("left join patients as p on p.hid = h.hid").
		Group("hid, title").
		Order("count DESC").
		Find(&count).Error; err != nil {
		return
	}
	return
}

/*
//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
*/
