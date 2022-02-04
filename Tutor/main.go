package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-playground/validator/v10"

	"encoding/json"
	"fmt"

	"github.com/gorilla/handlers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func initializeRouter() {

	router := mux.NewRouter()
	// This is to allow the headers, origins and methods all to access CORS resource sharing
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	router.HandleFunc("/api/v1/tutor/GetAllTutor", GetAllTutor).Methods("GET")
	router.HandleFunc("/api/v1/tutor/GetaTutorByEmail/{email}", GetaTutorByEmail).Methods("GET")
	router.HandleFunc("/api/v1/tutor/GetaTutorById/{tutor_id}", GetaTutorById).Methods("GET")
	router.HandleFunc("/api/v1/tutor/CreateNewTutor", CreateNewTutor).Methods("POST")
	router.HandleFunc("/api/v1/tutor/UpdateTutorAccountByEmail/{email}", UpdateTutorAccountByEmail).Methods("PUT")
	router.HandleFunc("/api/v1/tutor/DeleteTutorAccountByEamil/{email}", DeleteTutorAccountByEamil).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9181", handlers.CORS(headers, origins, methods)(router)))
}

func main() {
	initialMigration()
	initializeRouter()

}

type Tutor struct {
	Deleted      gorm.DeletedAt
	TutorID      int    `jason:"tutor_id"gorm:"primaryKey"`
	FirstName    string `json:"firstname" validate:"required"`
	Lastname     string `json:"lastname" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Descriptions string `json:"descriptions" validate:"required"`
}

const ADB = "root:root@tcp(127.0.0.2:3306)/assignment2?charset=utf8mb4&parseTime=True&loc=Local"

func initialMigration() {
	DB, err = gorm.Open(mysql.Open(ADB), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("cant conenct to the Database Please check the coneections strings")
	}
	DB.AutoMigrate(&Tutor{})
}

//Here is a fucntions to create new Tutor
func CreateNewTutor(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tutor Tutor
	var dbtutor Tutor
	json.NewDecoder(router.Body).Decode(&tutor)
	//to validate the inpute must be string not empty
	validate := validator.New()
	err2 := validate.Struct(tutor)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	//Validate duplications of email address
	err := DB.Where("email = ?", tutor.Email).First(&dbtutor).Error
	fmt.Println("tutor: " + dbtutor.FirstName)
	fmt.Println(err)
	if err == nil {
		fmt.Fprintf(w, "  The email you enter is already registerd  ")
		return
	} else {
		fmt.Fprintf(w, "  The Tutor Account Successfuly Registerd")

	}

	//if pass all validation Create passenger
	DB.Create(&tutor)
	json.NewEncoder(w).Encode(tutor)
}

//Get all registered Tutor
func GetAllTutor(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tutor []Tutor
	DB.Find(&tutor)
	json.NewEncoder(w).Encode(tutor)
}

//Here is a fucntions to get registerd Tutor by email
func GetaTutorByEmail(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var tutor Tutor
	err := DB.Where("email = ?", params["email"]).First(&tutor).Error
	if err != nil {
		//if user is not found
		fmt.Fprintf(w, "The email you enter is not registered")
		return
	} else {
		json.NewEncoder(w).Encode(tutor)
	}
}

func GetaTutorById(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var tutor Tutor
	err := DB.Where("tutor_id = ?", params["tutor_id"]).First(&tutor).Error
	if err != nil {
		//if user is not found
		fmt.Fprintf(w, "The email you enter is not registered")
		return
	} else {
		json.NewEncoder(w).Encode(tutor)
	}
}

//Here is a fucntions to update passenger
func UpdateTutorAccountByEmail(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var tutor Tutor
	err := DB.Where("email = ?", params["email"]).First(&tutor).Error

	if err != nil {
		fmt.Fprintf(w, "  The email you enter is not registerd  ")
		return
	} else {
		var tutor Tutor
		json.NewDecoder(router.Body).Decode(&tutor)
		DB.Model(&Tutor{}).Where("email=?", params["email"]).Updates(tutor)
		fmt.Fprintf(w, "Successfully update your account  ")

	}
}

//Here is the fucntions to delete Passenger

func DeleteTutorAccountByEamil(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var tutor Tutor

	err := DB.Where("email = ?", params["email"]).First(&tutor).Error

	if err != nil {
		fmt.Fprintf(w, "  The email you enter is not registerd  ")

	} else {

		json.NewDecoder(router.Body).Decode(&tutor)
		DB.Model(&Tutor{}).Where("email=?", params["email"]).Delete(&tutor)
		fmt.Fprintf(w, "Successfully Delete your account  ")

	}
}
