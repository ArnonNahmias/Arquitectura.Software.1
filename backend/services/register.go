package services


import (
    "backend/clients"
    "errors"


    "golang.org/x/crypto/bcrypt"
)


func RegisterS(username, password string, tipo string) error {
    // Check if the user already exists
    err := clients.SearchUser(username)
    if err == nil {
        return errors.New("username already taken")
    }


    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return errors.New("failed to hash password")
    }

    
    // Create the new user
    err = clients.CreateUser(username, string(hashedPassword), tipo)
    if err != nil {
        return errors.New("failed to create user")
    }


    return nil
}
