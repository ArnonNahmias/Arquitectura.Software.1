package main

import (
	"log"
)

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// getAllCourses retrieves all courses from the database
func getAllCourses() ([]Course, error) {
	rows, err := db.Query("SELECT id, name, description FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

// saveCourse saves a new course to the database
func saveCourse(course Course) error {
	stmt, err := db.Prepare("INSERT INTO courses(name, description) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Name, course.Description)
	if err != nil {
		return err
	}

	log.Println("New course added successfully")
	return nil
}
