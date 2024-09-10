package model

import "time"

type User struct {
	Id             string    `json:"id" gorm:"primaryKey;type:text" redis:"id"`
	Username       string    `json:"username" gorm:"type:varchar(25);uniqueIndex;not null" redis:"username"`
	Email          string    `json:"email" gorm:"type:varchar(25);uniqueIndex;not null" redis:"email"`
	Password       string    `json:"password,omitempty" gorm:"type:varchar(25);not null" redis:"-"`
	FollowerCount  int       `json:"follower_count" gorm:"type:int;default:0" redis:"-"`
	FollowingCount int       `json:"following_count" gorm:"type:int;default:0" redis:"-"`
	Salt           string    `json:"-" gorm:"type:text" redis:"-"`
	Token          string    `json:"-" gorm:"type:text" redis:"-"`
	CreatedOn      time.Time `json:"created_on" gorm:"type:timestamp"`
	EncodedData    string    `json:"-" redis:"data_created" gorm:"-"`
	Followers      []string  `json:"followers" gorm:"-"`
	Following      []string  `json:"following" gorm:"-"`
}

type UserFollowers struct {
	Follower string `json:"follower" gorm:"primaryKey;type:varchar(55) REFERENCES users(username); constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Followee string `json:"followee" gorm:"primaryKey;type:varchar(55) REFERENCES users(username); constraint:OnUpdate:CASCADE;not null"`
}
