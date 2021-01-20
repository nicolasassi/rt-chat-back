package operations

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RoomEntity struct {
	Name         string    `json:"name" bson:"name"`
	CreationDate time.Time `json:"creation_date" bson:"creation_date"`
}

type Rooms struct {
	Col *mongo.Collection
}

func NewRooms(client *mongo.Client, dbName, colName string) *Rooms {
	return &Rooms{
		Col: client.Database(dbName).Collection(colName),
	}
}

func (r Rooms) InsertNewRoom(ctx context.Context, name string) error {
	_, err := r.Col.InsertOne(ctx, RoomEntity{Name: name, CreationDate: time.Now()})
	if err != nil {
		return err
	}
	return nil
}

func (r Rooms) GetRoomByName(ctx context.Context, name string) (*RoomEntity, error) {
	roomEntity := new(RoomEntity)
	if err := r.Col.FindOne(ctx, bson.M{"name": name}).Decode(roomEntity); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return roomEntity, nil
}

func (r Rooms) GetRooms(ctx context.Context) (*[]RoomEntity, error) {
	c, err := r.Col.Find(ctx, bson.D{})
	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, nil
		}
		return nil, err
	}
	var roomEntities []RoomEntity
	if err := c.All(ctx, &roomEntities); err != nil {
		return nil, err
	}
	return &roomEntities, nil
}

func (r Rooms) DeleteRoomByName(ctx context.Context, name string) error {
	if err := r.Col.FindOneAndDelete(ctx, bson.M{"name": name}).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}
	return nil
}
