package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Make course file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Make fake DB
var courses []Course

func main() {
	fmt.Println("API Request")
	r := mux.NewRouter()

	// Seeding

	courses = append(courses, Course{CourseID: "2", CourseName: "Full-Stack", CoursePrice: "349", Author: &Author{FullName: "Ankit Pandey", Website: "udemy.com"}})
	courses = append(courses, Course{CourseID: "3", CourseName: "Devops", CoursePrice: "499", Author: &Author{FullName: "Vikash Kumar", Website: "udemy.com"}})

	//Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listening to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// controler files
// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop find there matching and return the response

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// for empty {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Send some data.... Empty list")
		return
	}

	//generate unique ID
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[:index+1]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Send some data into the list.")
		return
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete the course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[:index+1]...)
			break
		}
	}
}
