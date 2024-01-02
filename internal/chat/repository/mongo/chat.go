package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Messages struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Text              string             `bson:"text"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
	Author_user_id    int                `bson:"author_user_id"`
	Recipient_user_id int                `bson:"recipient_user_id"`
	Is_read           bool               `bson:"is_read"`
}

type MessagesRepository struct {
	db *mongo.Collection
}

func NewMessagesRepository(db *mongo.Database, collection string) *MessagesRepository {
	return &MessagesRepository{
		db: db.Collection(collection),
	}
}

// func (r MessagesRepository) CreateUser(ctx context.Context, m *models.User) error {
// 	model := toMongoMessages(m)
// 	res, err := r.db.InsertOne(ctx, model)
// 	if err != nil {
// 		return err
// 	}

// 	m.ID = res.InsertedID.(primitive.ObjectID).Hex()
// 	return nil
// }

// func toMongoMessages(m *models.Messages) *Messages {
// 	return &Messages{
// 		Text:              m.Text,
// 		CreatedAt:         m.CreatedAt,
// 		UpdatedAt:         m.UpdatedAt,
// 		Author_user_id:    m.Author_user_id,
// 		Recipient_user_id: m.Recipient_user_id,
// 		Is_read:           m.Is_read,
// 	}
// }

// func (r MessagesRepository) GetUser(ctx context.Context, , password string) (*models.User, error) {
// 	user := new(Messages)
// 	err := r.db.FindOne(ctx, bson.M{
// 		"username": username,
// 		"password": password,
// 	}).Decode(user)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return toModel(), nil
// }

// func toModel(m *Messages) *models.Messages {
// 	return &models.Messages{
// 		Id:                m.ID.Hex(),
// 		Text:              m.Text,
// 		CreatedAt:         m.CreatedAt,
// 		UpdatedAt:         m.UpdatedAt,
// 		Author_user_id:    m.Author_user_id,
// 		Recipient_user_id: m.Recipient_user_id,
// 		Is_read:           m.Is_read,
// 	}
// }
