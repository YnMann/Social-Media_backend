package mongo

import (
	"context"

	"github.com/YnMann/chat_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db *mongo.Collection
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

func NewAuthRepository(db *mongo.Database, collection string) *AuthRepository {
	return &AuthRepository{
		db: db.Collection(collection),
	}
}

func (r AuthRepository) CreateUser(ctx context.Context, user *models.User) error {
	model := toMongoUser(user)
	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func toMongoUser(u *models.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

func (r AuthRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := new(User)
	err := r.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func toModel(u *User) *models.User {
	return &models.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
	}
}
