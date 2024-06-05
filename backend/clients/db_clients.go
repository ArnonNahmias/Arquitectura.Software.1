package clients

import (
	"backend/dao"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing database...")
	dsn := "root:admin@tcp(127.0.0.1:3306)/proyecto?charset=utf8mb3&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	log.Println("Database connected successfully")

	Migrate()
	SeedDB()
}

func Migrate() {
	log.Println("Migrating database...")
	err := DB.AutoMigrate(&dao.Course{}, &dao.Usuario{}, &dao.Subscription{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully")
}

func SeedDB() {
	log.Println("Seeding database...")
	cursos := []dao.Course{
		{Nombre: "Ingles B2", Dificultad: "Medio", Precio: 45, Direccion: "José Roque Funes 1511 X5000ABE Córdoba"},
		{Nombre: "Hacking Etico", Dificultad: "Dificil", Precio: 99, Direccion: "Las Delicias"},
		{Nombre: "Marketing Digital", Dificultad: "Medio", Precio: 55, Direccion: "Rio Grande"},
		{Nombre: "C++", Dificultad: "Medio", Precio: 79, Direccion: "Cañuelas Country Golf"},
	}

	for _, curso := range cursos {
		DB.FirstOrCreate(&curso, dao.Course{Nombre: curso.Nombre})
	}

	usuarios := []dao.Usuario{
		{NombreUsuario: "ArnonNahmias", Contrasena: hashPassword("Arnon123")},
		{NombreUsuario: "Justi8", Contrasena: hashPassword("JUsti02")},
		{NombreUsuario: "Joako", Contrasena: hashPassword("RiverPLate")},
		{NombreUsuario: "Felipe08", Contrasena: hashPassword("Felipe05")},
	}

	for _, usuario := range usuarios {
		DB.FirstOrCreate(&usuario, dao.Usuario{NombreUsuario: usuario.NombreUsuario})
	}

	suscripciones := []dao.Subscription{
		{UserID: 1, CourseID: 1},
		{UserID: 2, CourseID: 2},
		{UserID: 3, CourseID: 3},
		{UserID: 4, CourseID: 4},
	}

	for _, suscripcion := range suscripciones {
		DB.FirstOrCreate(&suscripcion, dao.Subscription{UserID: suscripcion.UserID, CourseID: suscripcion.CourseID})
	}
	log.Println("Database seeded successfully")
}

func hashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SelectCoursesWithFilter(query string) ([]dao.Course, error) {
	var courses []dao.Course
	result := DB.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func SelectUserByEmail(email string) (dao.IdUsuario, error) {
	var user dao.IdUsuario
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return dao.IdUsuario{}, fmt.Errorf("not found user with email: %s", email)
	}
	return user, nil
}

func SelectUserByID(id int64) (dao.IdUsuario, error) {
	var user dao.IdUsuario
	result := DB.First(&user, id)
	if result.Error != nil {
		return dao.IdUsuario{}, fmt.Errorf("not found user with ID: %d", id)
	}
	return user, nil
}

func InsertSubscription(userID int64, courseID int64) error {
	var subscription dao.Subscription
	result := DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&subscription)
	if result.Error == nil {
		return fmt.Errorf("user %d is already subscribed to course %d", userID, courseID)
	}

	subscription = dao.Subscription{
		UserID:       userID,
		CourseID:     courseID,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	}

	result = DB.Create(&subscription)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SelectCourseByID(id int64) (dao.Course, error) {
	var course dao.Course
	result := DB.First(&course, id)
	if result.Error != nil {
		return dao.Course{}, fmt.Errorf("not found course with ID: %d", id)
	}
	return course, nil
}
