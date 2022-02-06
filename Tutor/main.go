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

//fucntions to intiate routers
func initializeRouter() {

	router := mux.NewRouter()
	// This is to allow the headers, origins and methods all to access CORS resource sharing
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	router.HandleFunc("/api/v1/tutor/GetAllTutor", GetAllTutor).Methods("GET")                                        //get All Tutor
	router.HandleFunc("/api/v1/tutor/GetaTutorByEmail/{email}", GetaTutorByEmail).Methods("GET")                      // Get tutor by email
	router.HandleFunc("/api/v1/tutor/GetaTutorById/{tutor_id}", GetaTutorById).Methods("GET")                         // Get tutor by name
	router.HandleFunc("/api/v1/tutor/CreateNewTutor", CreateNewTutor).Methods("POST")                                 // Create new tutor
	router.HandleFunc("/api/v1/tutor/UpdateTutorAccountByEmail/{email}", UpdateTutorAccountByEmail).Methods("PUT")    //Edit Tutor by email
	router.HandleFunc("/api/v1/tutor/DeleteTutorAccountByEamil/{email}", DeleteTutorAccountByEamil).Methods("DELETE") // Delete Tutor by email
	fmt.Println("Listening to port 9181 ")
	log.Fatal(http.ListenAndServe(":9181", handlers.CORS(headers, origins, methods)(router)))
}

//this fuctions to start both initiate migrations and routers
func main() {
	initialMigration()
	initializeRouter()

}

//tutor stuct
type Tutor struct {
	Deleted      gorm.DeletedAt
	TutorID      int    `json:"tutor_id"gorm:"primaryKey"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Descriptions string `json:"descriptions" validate:"required"`
}

//Database connections strings
const ADB = "root:root@tcp(db:3306)/assignment2?charset=utf8mb4&parseTime=True&loc=Local"

//fucntions to intiate migrations
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
	fmt.Println("tutor: " + dbtutor.Name)
	fmt.Println(err)
	if err == nil {
		fmt.Fprintf(w, "  The email you enter is already registerd  ")
		w.WriteHeader(http.StatusIMUsed)
		return
	} else {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "  The Tutor Account Successfuly Registerd")
		w.WriteHeader(http.StatusAccepted)

	}

	//if pass all validation Create passenger
	DB.Create(&tutor)
	json.NewEncoder(w).Encode(tutor)
}

//Get all registered Tutor
func GetAllTutor(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// make a variable list of tutor
	var tutor []Tutor
	//find the list of tutor from database
	DB.Find(&tutor)
	//encode and decode the data
	json.NewEncoder(w).Encode(tutor)
}

//Here is a fucntions to get registerd Tutor by email
func GetaTutorByEmail(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//mux router parameter
	params := mux.Vars(router)
	// make a variable of tutor
	var tutor Tutor
	//validate the email address to database
	err := DB.Where("email = ?", params["email"]).First(&tutor).Error
	if err != nil {
		//error message if email is not registered
		fmt.Fprintf(w, "The email you enter is not registered")
		return
	} else {
		//encode and decode the data
		json.NewEncoder(w).Encode(tutor)
	}
}

//get tutor by id
func GetaTutorById(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var tutor Tutor
	//validate the id if is register to the database.
	err := DB.Where("tutor_id = ?", params["tutor_id"]).First(&tutor).Error
	if err != nil {
		//error message if email is not registered
		fmt.Fprintf(w, "The Tutor ID you enter is not registered")
		return
	} else {
		//decode the data
		json.NewEncoder(w).Encode(tutor)
	}
}

//Here is a fucntions to update passenger
func UpdateTutorAccountByEmail(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//mux router parameter
	params := mux.Vars(router)
	// make a variable of tutor
	var tutor Tutor
	//validate the email if is register to the database.
	err := DB.Where("email = ?", params["email"]).First(&tutor).Error

	if err != nil {
		//error message if email is not registered
		fmt.Fprintf(w, "  The email you enter is not registerd  ")
		return
	} else {
		var tutor Tutor
		//decode the data
		json.NewDecoder(router.Body).Decode(&tutor)
		//making changes to the database base on the email selected
		DB.Model(&Tutor{}).Where("email=?", params["email"]).Updates(tutor)
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Successfully update your account  ")

	}
}

//Here is the fucntions to delete Passenger
func DeleteTutorAccountByEamil(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//mux router parameter
	params := mux.Vars(router)
	// make a variable of tutor
	var tutor Tutor

	err := DB.Where("email = ?", params["email"]).First(&tutor).Error
	//error message if email is not registered Tutor
	if err != nil {
		fmt.Fprintf(w, "  The email you enter is not registerd  ")

	} else {
		//decode the data
		json.NewDecoder(router.Body).Decode(&tutor)
		DB.Model(&Tutor{}).Where("email=?", params["email"]).Delete(&tutor)

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Successfully Delete your account  ")

	}
}
