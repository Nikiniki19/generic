package service

import (
	"context"
	"errors"
	"generic/database"
	"generic/models"
	"generic/proto"
	"log"

	"gorm.io/gorm"
)

type FetchUser struct {
	proto.GenericRequestServer
}

func (f *FetchUser) FetchUser(ctx context.Context, userID *proto.GenericClientID) (*proto.GenericResponse, error) {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}

	useradata := models.GenericUsers{
		Model: gorm.Model{
			ID: uint(userID.Id),
		},
	}
	result := db.Where("id = ?", useradata.ID).First(&useradata)
	if result.Error != nil && result.RowsAffected == 0 {
		return &proto.GenericResponse{}, errors.New("could not fetch the user")
	}
	return &proto.GenericResponse{
		Username: useradata.Username,
		Email:    useradata.Email,
	}, nil
}
