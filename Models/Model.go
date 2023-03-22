package Models

func (b *Hospitals) TableName() string {
	return "hospitals"
}

type Hospitals struct {
	Hid   string `json:"hid"`
	Title string `json:"title"`
}

type Patients struct {
	Hnid      string `json:"hnid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Hid       string `json:"hid"`
}

type PatientCovidStatus struct {
	Hnid        string `json:"hnid"`
	CovidStatus string `json:"covid_status"`
}

type AllCovidPatientInfo struct {
	Hnid        string `json:"hnid"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CovidStatus string `json:"covid_status"`
}

type CountPatient struct {
	Count int    `json:"count"`
	Hid   string `json:"hid"`
	Title string `json:"title"`
}
