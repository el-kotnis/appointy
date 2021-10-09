package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
	Post     *Post         `json:"posts" bson:"posts"`
}

type Post struct {
	Id               bson.ObjectId       `json:"id" bson:"id"`
	Caption          string              `json:"caption" bson:"caption"`
	Image_URL        string              `json:"img_url" bson:"img_url"`
	Posted_Timestamp bson.MongoTimestamp `json:"timestamp" bson:"timestamp"`
}
