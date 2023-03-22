CREATE TABLE hospitals (
    hid varchar(255) PRIMARY KEY,
    title varchar(255)
)

CREATE TABLE patients (
    hnid varchar(255) PRIMARY KEY,
    first_name varchar(255),
    last_name varchar(255),
    hid varchar(255),
    CONSTRAINT FK_hid FOREIGN KEY (hid) REFERENCES hospitals(hid)
)

CREATE TABLE patient_covid_status (
    hnid varchar(255) PRIMARY KEY,
    covid_status varchar(255),
    CONSTRAINT FK_hnid FOREIGN KEY (hnid) REFERENCES patients(hnid)

)

INSERT INTO hospitals VALUE ("0001", "St. Violet Hospital")
INSERT INTO hospitals VALUE ("0002", "The Unity Health Center")
INSERT INTO hospitals VALUE ("0003", "John Medical Center")

INSERT INTO patients VALUE ("0001", "Ayesha", "Jenkins", "0001")
INSERT INTO patients VALUE ("0002", "Savannah", "Mcintosh", "0002")
INSERT INTO patients VALUE ("0003", "Harley", "Greene", "0002")
INSERT INTO patients VALUE ("0004", "Tilly", "Frazier", "0001")
INSERT INTO patients VALUE ("0005", "Gloria", "Hull", "0003")

INSERT INTO patient_covid_status VALUE ("0001", "Positive")
INSERT INTO patient_covid_status VALUE ("0002", "Positive")
INSERT INTO patient_covid_status VALUE ("0003", "Negative")
INSERT INTO patient_covid_status VALUE ("0004", "Positive")
INSERT INTO patient_covid_status VALUE ("0005", "Negative")
