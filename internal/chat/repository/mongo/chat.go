package mongo

import (
	"context"
	"time"

	"github.com/YnMann/chat_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Messages struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Sender    string             `bson:"sender"`
	Recipient string             `bson:"recipient"`
	Content   string             `bson:"content"`
	ServerIP  string             `bson:"server_ip"`
	SenderIP  string             `bson:"sender_ip"`
	IsRead    bool               `bson:"is_read"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type ChatRepository struct {
	mdb *mongo.Collection
	udb *mongo.Collection
}

func NewMessagesRepository(
	db *mongo.Database,
	msg_collection string,
	u_collection string,
) *ChatRepository {
	return &ChatRepository{
		mdb: db.Collection(msg_collection),
		udb: db.Collection(u_collection),
	}
}

func (r ChatRepository) SetUserOnlineStatus(ctx context.Context, userID string, isOnline bool) error {
	filter := bson.M{"_id": userID}

	update := bson.M{"$set": bson.M{"is_online": isOnline}}

	_, err := r.udb.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// CreateMsg creates a new message in MongoDB.
func (c ChatRepository) CreateMsg(ctx context.Context, m *models.Messages) error {
	// Convert the models.Messages to a MongoDB document
	model := toMongoMsg(m)

	// Insert the document into MongoDB
	res, err := c.mdb.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	// Set the ID of the model from the inserted document
	m.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func toMongoMsg(m *models.Messages) *Messages {
	return &Messages{
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Content:   m.Content,
		ServerIP:  m.ServerIP,
		SenderIP:  m.SenderIP,
		IsRead:    m.IsRead,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (r ChatRepository) GetMsg(
	ctx context.Context,
	sender string,
	sender_ip string,
	recipient string,
) (*models.Messages, error) {
	msg := new(Messages)
	err := r.mdb.FindOne(ctx, bson.M{
		"sender":    sender,
		"sender_ip": sender_ip,
		"recipient": recipient,
	}).Decode(msg)

	if err != nil {
		return nil, err
	}

	return toModelMsg(msg), nil
}

func toModelMsg(m *Messages) *models.Messages {
	return &models.Messages{
		ID:        m.ID.Hex(),
		Sender:    m.Sender,
		Recipient: m.Recipient,
		Content:   m.Content,
		ServerIP:  m.ServerIP,
		SenderIP:  m.SenderIP,
		IsRead:    m.IsRead,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
