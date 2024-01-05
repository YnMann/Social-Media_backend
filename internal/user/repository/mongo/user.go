package mongo

import (
	"context"
	"time"

	"github.com/YnMann/chat_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	auth "github.com/YnMann/chat_backend/internal/auth/repository/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

type UserProfile struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Username  string             `bson:"username"`
	PhotoURL  string             `bson:"photo_url"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	IsOnline  bool               `bson:"is_online"`
	IsBanned  bool               `bson:"is_banned"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewUserRepository(db *mongo.Database, collection string) (*UserRepository, *auth.AuthRepository) {
	userRepo := &UserRepository{
		db: db.Collection(collection),
	}

	// Creating an autorepo inside the New User Repository, passing the same collection to it
	authRepo := auth.NewAuthRepository(db, collection)

	return userRepo, authRepo
}

func (r UserRepository) GetUserProfile(ctx context.Context, uID string) (*models.User, error) {
	user := new(UserProfile)

	err := r.db.FindOne(ctx, bson.M{"_id": uID}).Decode(user)
	if err != nil {
		return nil, err
	}

	return toModelProfile(user), nil
}

func toModelProfile(u *UserProfile) *models.User {
	return &models.User{
		ID:        u.ID.Hex(),
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		PhotoURL:  u.PhotoURL,
		IsBanned:  u.IsBanned,
	}
}

type Contacts struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	PhotoURL  string             `bson:"photo_url"`
}

func (r UserRepository) GetContacts(ctx context.Context) ([]*models.Contacts, error) {
	var contacts []*Contacts
	// The filter is empty to get all users
	filter := bson.M{}

	// A projection indicating which fields to select
	projection := bson.M{
		"photo_url":  1,
		"id":         1,
		"first_name": 1,
		"last_name":  1,
	}

	cursor, err := r.db.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var contact Contacts
		if err := cursor.Decode(&contact); err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toModelContacts(contacts), nil
}

func toModelContacts(c []*Contacts) []*models.Contacts {
	contacts := make([]*models.Contacts, len(c))

	for i, user := range c {
		contacts[i] = &models.Contacts{
			ID:        user.ID.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			PhotoURL:  user.PhotoURL,
		}
	}

	return contacts
}
