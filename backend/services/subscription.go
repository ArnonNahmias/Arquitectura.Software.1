package services

import (
	"errors"
	"log"
	"time"

	"backend/dao"
	"backend/domain"
	"gorm.io/gorm"
)

// GetSubscriptions retrieves all subscriptions.
func GetSubscriptions() ([]domain.Subscription, error) {
	var subscriptions []domain.Subscription
	if err := dao.DB.Find(&subscriptions).Error; err != nil {
		log.Printf("Error fetching subscriptions: %v", err)
		return nil, err
	}
	return subscriptions, nil
}

// CreateSubscription creates a new subscription.
func CreateSubscription(subscription domain.Subscription) (domain.Subscription, error) {
	// Check if the user exists
	log.Printf("Checking if user exists: %d", subscription.IdUsuario)
	if err := checkUserExists(subscription.IdUsuario); err != nil {
		log.Printf("Error checking user existence: %v", err)
		return subscription, err
	}

	// Check if the course exists
	log.Printf("Checking if course exists: %d", subscription.IdCurso)
	if err := checkCourseExists(subscription.IdCurso); err != nil {
		log.Printf("Error checking course existence: %v", err)
		return subscription, err
	}

	// Create the subscription
	subscription.CreatedAt = time.Now()
	subscription.UpdatedAt = time.Now()
	log.Println("Creating subscription in database...")
	if err := dao.DB.Create(&subscription).Error; err != nil {
		log.Printf("Error creating subscription: %v", err)
		return subscription, err
	}
	log.Println("Subscription created successfully in database")
	return subscription, nil
}

// DeleteSubscription deletes a subscription by ID.
func DeleteSubscription(id int64) error {
	var subscription domain.Subscription
	log.Printf("Fetching subscription by ID: %d", id)
	if err := dao.DB.First(&subscription, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Subscription not found: %v", err)
			return errors.New("subscription not found")
		}
		log.Printf("Error fetching subscription: %v", err)
		return err
	}
	log.Println("Deleting subscription from database...")
	if err := dao.DB.Delete(&subscription).Error; err != nil {
		log.Printf("Error deleting subscription: %v", err)
		return err
	}
	log.Println("Subscription deleted successfully from database")
	return nil
}

// checkUserExists checks if a user exists by ID.
func checkUserExists(userID int64) error {
	var user domain.Usuario
	log.Printf("Fetching user by ID: %d", userID)
	if err := dao.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("User not found: %v", err)
			return errors.New("user not found")
		}
		log.Printf("Error fetching user: %v", err)
		return err
	}
	log.Println("User exists")
	return nil
}

// checkCourseExists checks if a course exists by ID.
func checkCourseExists(courseID int64) error {
	var course domain.Course
	log.Printf("Fetching course by ID: %d", courseID)
	if err := dao.DB.First(&course, courseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Course not found: %v", err)
			return errors.New("course not found")
		}
		log.Printf("Error fetching course: %v", err)
		return err
	}
	log.Println("Course exists")
	return nil
}
