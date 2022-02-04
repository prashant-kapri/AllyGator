package controller

import (
	"database/sql"

	"allygator.com/gatorweb/models"
)

type User struct {
	StudentId      int    `json:"idStudent"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	Department     int    `json:"department"`
	Password       string `json:"password"`
	UFmail         string `json:"email"`
	Gender         string `json:"gender"`
	Course         string `json:"course"`
	URL            string `json:"url"`
	Nationality    string `json:"nationality"`
	Profile        string `json:"profile"`
	Specialization string `json:"specialization"`
}

//This function retrieves the list of all the students from the database
func GetUsers() ([]User, error) {

	rows, err := models.DB.Query("SELECT idStudent, firstname, lastname, department, email, gender, course, url, nationality, profile, specialization from users")
	// Select * from user TABLE
	if err != nil {
		return nil, err
	}
	//Checking for an error on the query and defering the Close() method for the tuples to ensure it doesn’t stay open after we’re done.
	defer rows.Close()

	people := make([]User, 0)

	for rows.Next() {
		singleUser := User{}
		err = rows.Scan(&singleUser.StudentId, &singleUser.FirstName, &singleUser.LastName, &singleUser.Department, &singleUser.UFmail, &singleUser.Gender, &singleUser.Course, &singleUser.URL, &singleUser.Nationality, &singleUser.Profile, &singleUser.Specialization)

		if err != nil {
			return nil, err
		}
		people = append(people, singleUser)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return people, err
}

//This function is used to append the student details into the database
func AddUsers(newUser User) (bool, error) {

	tx, err := models.DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO users (firstname, lastname, department, password, email, gender, course, url, nationality, profile, specialization) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newUser.FirstName, newUser.LastName, newUser.Department, newUser.Password, newUser.UFmail, newUser.Gender, newUser.Course, newUser.URL, newUser.Nationality, newUser.Profile, newUser.Specialization)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

//This function is used to retrieve the student details by ID
func GetUserById(idStudent string) (User, error) {

	stmt, err := models.DB.Prepare("SELECT idStudent, firstname, lastname, department, password, email, gender, course, url, nationality, profile, specialization from users WHERE idStudent = ?")

	if err != nil {
		return User{}, err
	}

	user := User{}

	sqlErr := stmt.QueryRow(idStudent).Scan(&user.StudentId, &user.FirstName, &user.LastName, &user.Department, &user.Password, &user.UFmail, &user.Gender, &user.Course, &user.URL, &user.Nationality, &user.Profile, &user.Specialization)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return User{}, nil
		}
		return User{}, sqlErr
	}
	return user, nil
}
