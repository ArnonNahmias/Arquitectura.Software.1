package services

import (
	"backend/clients"
	"backend/domain"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

/*type Course struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}*/

type CourseService struct {
	DB *gorm.DB
}

func Search(query string) ([]domain.Course, error) {
	trimmed := strings.TrimSpace(query)

	courses, err := clients.SelectCoursesWithFilterName(trimmed)
	if err != nil {
		return nil, fmt.Errorf("error getting courses from DB: %w", err)
	}

	results := make([]domain.Course, 0)
	for _, course := range courses {
		results = append(results, domain.Course{
			IdCurso:    int64(course.IdCurso), // Convert int to int64
			Nombre:     course.Nombre,
			Dificultad: course.Dificultad,
			Precio:     strconv.Itoa(course.Precio), // Convert int to string
			ImageURL:   course.ImageURL,
			CreatedAt:  course.CreatedAt,
			UpdatedAt:  course.UpdatedAt,
		})
	}
	return results, nil
}

func SearchbyID(query int) ([]domain.Course, error) {
	// No es necesario hacer TrimSpace en un entero

	// Llamamos a la función para seleccionar cursos con el filtro (el ID en este caso)
	courses, err := clients.SelectCoursesWithFilterID(query)
	if err != nil {
		return nil, fmt.Errorf("error getting courses from DB: %w", err)
	}

	// Creamos un slice para los resultados
	results := make([]domain.Course, 0)
	for _, course := range courses {
		// Añadimos cada curso al slice de resultados, haciendo las conversiones necesarias
		results = append(results, domain.Course{
			IdCurso:    int64(course.IdCurso), // Convertimos int a int64
			Nombre:     course.Nombre,
			Dificultad: course.Dificultad,
			Precio:     strconv.Itoa(course.Precio), // Convertimos int a string
			ImageURL:   course.ImageURL,
			CreatedAt:  course.CreatedAt,
			UpdatedAt:  course.UpdatedAt,
		})
	}

	// Devolvemos los resultados
	return results, nil
}



/*func Subscribe(userID int64, courseID int64) error {
	if _, err := clients.SelectUserByID(userID); err != nil {
		return fmt.Errorf("error getting user from DB: %w", err)
	}

	if _, err := clients.SelectCourseByID(courseID); err != nil {
		return fmt.Errorf("error getting course from DB: %w", err)
	}

	if err := clients.InsertSubscription(userID, courseID); err != nil {
		return fmt.Errorf("error creating subscription in DB: %w", err)
	}

	return nil
}
*/
