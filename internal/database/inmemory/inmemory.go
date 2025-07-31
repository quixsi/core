package inmemory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/quixsi/app/internal/database"
	"github.com/quixsi/app/internal/models"
)

type InMemoryDB struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*models.User
}

func New() database.Database {
	now := time.Now()
	id1 := uuid.New()
	id2 := uuid.New()
	
	users := map[uuid.UUID]*models.User{
		id1: {
			ID:        id1,
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		id2: {
			ID:        id2,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			CreatedAt: &now,
			UpdatedAt: &now,
		},
	}

	return &InMemoryDB{
		users: users,
	}
}

func (db *InMemoryDB) GetUsers(ctx context.Context) ([]*models.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	users := make([]*models.User, 0, len(db.users))
	for _, user := range db.users {
		users = append(users, user)
	}
	return users, nil
}

func (db *InMemoryDB) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	user, exists := db.users[id]
	if !exists {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}

func (db *InMemoryDB) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	now := time.Now()
	newUser := &models.User{
		ID:        uuid.New(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	db.users[newUser.ID] = newUser
	return newUser, nil
}

func (db *InMemoryDB) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	existingUser, exists := db.users[user.ID]
	if !exists {
		return nil, fmt.Errorf("user with id %s not found", user.ID)
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	now := time.Now()
	existingUser.UpdatedAt = &now

	return existingUser, nil
}

func (db *InMemoryDB) DeleteUser(ctx context.Context, id uuid.UUID) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.users[id]; !exists {
		return fmt.Errorf("user with id %s not found", id)
	}

	delete(db.users, id)
	return nil
}