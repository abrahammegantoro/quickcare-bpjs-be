package repositories

import (
	"context"
	"fmt"
	"time"

	// "github.com/alextanhongpin/go-mongo/models"
	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context, filter bson.M) (*models.User, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	var users *models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("error decoding user: %w", err)
	}

	return users, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, filter bson.M) (*models.User, error) {
	var user *models.User
	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, fmt.Errorf("error decoding user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, filter bson.M) error {
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}