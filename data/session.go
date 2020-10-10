package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/a-berahman/auth-api/common"
	"github.com/a-berahman/auth-api/config"
	"github.com/a-berahman/auth-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateSession returns session unique id
func CreateSession(uid string) (string, error) {
	if uid == "" {
		return "", fmt.Errorf("uid mustn't be empty")
	}
	mu := &sync.RWMutex{}
	mu.Lock()
	defer mu.Unlock()

	var wg sync.WaitGroup
	wg.Add(1)
	go checkLimitation(uid, &wg)

	// prepare and inserting session
	oid, err := insertSession(uid)
	if err != nil {
		return "", err
	}
	wg.Wait()
	//Check session count limitation and remove extera session

	return oid.Hex(), nil

}

func insertSession(uid string) (primitive.ObjectID, error) {

	s := &model.Session{
		ID:      primitive.NewObjectID(),
		Created: time.Now(),
		Status:  true,
		UserID:  uid,
	}

	session, err := config.Sessions.InsertOne(context.Background(), &s)
	if err != nil {
		return [12]byte{}, fmt.Errorf("face on error while insert session : %w", err)
	}

	oid, ok := session.InsertedID.(primitive.ObjectID)
	if !ok {
		return [12]byte{}, fmt.Errorf("Can not convert to OID")
	}
	return oid, nil
}
func checkLimitation(uid string, wg *sync.WaitGroup) {
	defer wg.Done()
	isExceeded, sessionList, err := isLimitExceeded(uid)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if isExceeded {
		deleteExceededSession(sessionList)
	}

}

// isLimitExceeded checks session exceed and return sessions that should remove
func isLimitExceeded(uid string) (bool, []model.Session, error) {

	filter := bson.M{"userid": uid}

	opts := options.Find()
	opts.SetSort(bson.M{"Created": -1})
	sessionList := []model.Session{}

	cur, err := config.Sessions.Find(context.Background(), filter, opts)
	if err != nil {
		return true, nil, fmt.Errorf("face on error while finding session by UserID : %w", err)
	}
	defer cur.Close(context.Background())

	err = cur.All(context.Background(), &sessionList)
	if cur != nil && err != nil {
		return true, nil, fmt.Errorf("face on error while decoding sessions : %w", err)
	}
	length := len(sessionList)

	lc, _ := strconv.Atoi(os.Getenv(common.LIMITATION_COUNT))
	if length > lc {

		return true, sessionList[:length-lc], nil
	}
	return false, nil, nil
}

func deleteExceededSession(sessionList []model.Session) {

	oids := make([]primitive.ObjectID, len(sessionList))

	for i, v := range sessionList {
		oids[i] = v.ID
	}
	filter := bson.M{"_id": bson.M{"$in": oids}}

	_, err := config.Sessions.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatalln("while delete exceed sessions face on : ", err)
	}

	// if err := config.Sessions.Drop(ctx); err != nil {
	// 	log.Fatalln("while drop session delete context face on : ", err)
	// }

}
