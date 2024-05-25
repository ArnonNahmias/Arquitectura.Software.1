package domain

import "time"

// Course representa un curso en el sistema.
type Course struct {
	ID           int64     `json:"id"`            // ID del curso
	Title        string    `json:"title"`         // Título del curso
	Description  string    `json:"description"`   // Descripción del curso
	Category     string    `json:"category"`      // Categoría del curso (por definir)
	CreationDate time.Time `json:"creation_date"` // Fecha de creación del curso
	LastUpdated  time.Time `json:"last_updated"`  // Fecha de última actualización del curso
}

// SearchRequest representa una solicitud de búsqueda de cursos.
type SearchRequest struct {
	Query string `json:"query"` // Consulta de búsqueda
}

// SearchResponse representa una respuesta de búsqueda de cursos.
type SearchResponse struct {
	Results []Course `json:"results"` // Resultados de la búsqueda
}

// SubscribeRequest representa una solicitud de suscripción a un curso.
type SubscribeRequest struct {
	UserID   int64 `json:"user_id"`   // ID del usuario
	CourseID int64 `json:"course_id"` // ID del curso
}
