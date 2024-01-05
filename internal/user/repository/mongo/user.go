package mongo

import (
	"context"
	"time"

	"github.com/YnMann/chat_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
