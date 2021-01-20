package operations

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ChatEntity struct {
	RoomID   primitive.ObjectID `json:"room_id" bson:"room_id"`
	SentAt   time.Time          `json:"sent_at" bson:"sent_at"`
	UserID   string             `json:"user_id" bson:"user_id"`
	UserName string             `json:"user_name" bson:"user_name"`
	Message  string             `json:"message" bson:"message"`
}

type Chats struct {
	Col *mongo.Collection
}

func NewChats(client *mongo.Client, dbName, colName string) *Rooms {
	return &Rooms{
		Col: client.Database(dbName).Collection(colName),
	}
}

func (r Rooms) InsertNewChat(ctx context.Context, chat ChatEntity) error {
	_, err := r.Col.InsertOne(ctx, chat)
	if err != nil {
		return err
	}
	return nil
}

func (r Rooms) GetChatsByRoomID(ctx context.Context, roomID primitive.ObjectID) (*[]ChatEntity, error) {
	c, err := r.Col.Find(ctx, bson.M{"room_id": roomID})
	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, nil
		}
		return nil, err
	}
	var chatEntities []ChatEntity
	if err := c.All(ctx, &chatEntities); err != nil {
		return nil, err
	}
	return &chatEntities, nil
}

func (r Rooms) GetChats(ctx context.Context) (*[]ChatEntity, error) {
	c, err := r.Col.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, nil
		}
		return nil, err
	}
	var chatEntities []ChatEntity
	if err := c.All(ctx, &chatEntities); err != nil {
		return nil, err
	}
	return &chatEntities, nil
}