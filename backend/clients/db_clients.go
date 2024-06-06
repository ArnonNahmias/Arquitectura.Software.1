package clients

import (
	"backend/dao"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing database...")
	dsn := "root:root@tcp(127.0.0.1:3306)/proyecto?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
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
		{Nombre: "Ingles B2", Dificultad: "Medio", Precio: 45, Direccion: "José Roque Funes 1511 X5000ABE Córdoba", ImageURL: "https://blog.hubspot.com/hs-fs/hubfs/freeonlinecourses-1.webp?width=595&height=400&name=freeonlinecourses-1.webp"},
	}
	for _, curso := range cursos {
		DB.Create(&curso)
	}
	log.Println("Database seeded successfully")
}

func SelectUser(username string) (dao.Usuario, error) {
	var user dao.Usuario
	result := DB.Where("Nombre_usuario = ?", username).Find(&user)
	if result.Error != nil {
		return dao.Usuario{}, result.Error
	}
	return user, nil
}

func SelectCoursesWithFilter(query string) ([]dao.Course, error) {
	var courses []dao.Course
	result := DB.Where("Nombre LIKE ? ", "%"+query+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}
