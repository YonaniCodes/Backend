package data

import (
	"a2sv-backend/task_manager/models"
	"errors"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	users      []models.User
	usersMutex sync.RWMutex
	nextUserID = 1
)

// RegisterUser creates a new user account
func RegisterUser(username, password string) (*models.User, error) {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Check if username already exists
	for _, user := range users {
		if user.Username == username {
			return nil, errors.New("username already exists")
		}
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Determine role: first user is admin, others are regular users
	role := "user"
	if len(users) == 0 {
		role = "admin"
	}

	// Create new user
	newUser := models.User{
		ID:       nextUserID,
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	users = append(users, newUser)
	nextUserID++

	return &newUser, nil
}

// AuthenticateUser validates user credentials and returns the user if valid
func AuthenticateUser(username, password string) (*models.User, error) {
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	// Find user by username
	var user *models.User
	for i := range users {
		if users[i].Username == username {
			user = &users[i]
			break
		}
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Return user without password
	userWithoutPassword := *user
	userWithoutPassword.Password = ""

	return &userWithoutPassword, nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*models.User, error) {
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	for i := range users {
		if users[i].Username == username {
			user := users[i]
			user.Password = "" // Don't return password
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

// PromoteUser promotes a user to admin role
func PromoteUser(username string) error {
	usersMutex.Lock()
	defer usersMutex.Unlock()

	for i := range users {
		if users[i].Username == username {
			users[i].Role = "admin"
			return nil
		}
	}

	return errors.New("user not found")
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*models.User, error) {
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	for i := range users {
		if users[i].ID == id {
			user := users[i]
			user.Password = "" // Don't return password
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

