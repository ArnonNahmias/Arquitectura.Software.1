package services

import (
    "errors"
    "log"
    "time"
    "backend/clients"
    "backend/dao"
    "gorm.io/gorm"
)

// CreateSubscription crea una nueva suscripción.
func CreateSubscription(subscription dao.Subscription) (dao.Subscription, error) {
    tx := clients.DB.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Verificar si el usuario existe
    var user dao.Usuario
    log.Printf("Checking if user exists: %d", subscription.IdUsuario)
    if err := tx.First(&user, subscription.IdUsuario).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            tx.Rollback()
            return subscription, errors.New("user not found")
        }
        tx.Rollback()
        return subscription, err
    }

    // Verificar si el curso existe
    var course dao.Course
    log.Printf("Checking if course exists: %d", subscription.IdCurso)
    if err := tx.First(&course, subscription.IdCurso).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            tx.Rollback()
            return subscription, errors.New("course not found")
        }
        tx.Rollback()
        return subscription, err
    }

    // Crear la suscripción
    subscription.CreatedAt = time.Now()
    subscription.UpdatedAt = time.Now()
    if err := tx.Create(&subscription).Error; err != nil {
        tx.Rollback()
        return subscription, err
    }

    if err := tx.Commit().Error; err != nil {
        return subscription, err
    }

    return subscription, nil
}


// GetUserSubscriptions fetches subscriptions for a given user ID.
func GetUserSubscriptions(userId int64) ([]dao.Course, error) {
	var subscriptions []dao.Subscription
	var courses []dao.Course

	// Fetch subscriptions for the user
	if err := clients.DB.Where("id_usuario = ?", userId).Find(&subscriptions).Error; err != nil {
		log.Printf("Error fetching subscriptions for user ID %d: %v", userId, err)
		return nil, err
	}

	// Fetch courses for each subscription
	for _, subscription := range subscriptions {
		var course dao.Course
		if err := clients.DB.First(&course, subscription.IdCurso).Error; err == nil {
			courses = append(courses, course)
		} else {
			log.Printf("Error fetching course ID %d for subscription ID %d: %v", subscription.IdCurso, subscription.IdSubscription, err)
		}
	}

	return courses, nil
}

// DeleteSubscription elimina una suscripción por ID.
func DeleteSubscription(id int) error {
	var subscription dao.Subscription
	if err := clients.DB.First(&subscription, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("subscription not found")
		}
		return err
	}
	if err := clients.DB.Delete(&subscription).Error; err != nil {
		return err
	}
	return nil
}
