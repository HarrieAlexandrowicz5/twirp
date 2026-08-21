package user

import (
	"context"
	"net/http"
	"github.com/twitchtv/twirp"
	"github.com/HarrieAlexandrowicz5/twirp/api/user"
)

// GetUserHandler handles the /user/get endpoint
func GetUserHandler(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	// Check if the user exists
	if req.UserId == "" {
		return nil, twirp.InvalidArgumentError("user_id", "user_id is required")
	}
	
	// Get the user data
	userData := getUserData(req.UserId)
	
	// Check if the user data is nil
	if userData == nil {
		return nil, twirp.NotFoundError("user_id", "user not found")
	}
	
	// Return the user data in JSON format
	return &user.GetUserResponse{
		User: &user.User{
			Id:   userData.Id,
			Name: userData.Name,
		},
	}, nil
}

// getUserData simulates getting user data from a database
func getUserData(userId string) *user.User {
	// Simulate getting user data from a database
	return &user.User{
		Id:   userId,
		Name: "John Doe",
	}
}