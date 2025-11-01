package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/Alexandru2984/twitter_clone/backend/config"
	"github.com/Alexandru2984/twitter_clone/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FollowUser(c *gin.Context) {
	targetUsername := c.Param("username")
	currentUserID := c.GetString("userID")

	currentObjID, err := primitive.ObjectIDFromHex(currentUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find target user
	var targetUser models.User
	err = collection.FindOne(ctx, bson.M{"username": targetUsername}).Decode(&targetUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Can't follow yourself
	if targetUser.ID == currentObjID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot follow yourself"})
		return
	}

	// Check if already following
	var currentUser models.User
	err = collection.FindOne(ctx, bson.M{"_id": currentObjID}).Decode(&currentUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Current user not found"})
		return
	}

	following := false
	for _, followID := range currentUser.Following {
		if followID == targetUser.ID {
			following = true
			break
		}
	}

	// Toggle follow
	if following {
		// Unfollow
		_, err = collection.UpdateOne(ctx, bson.M{"_id": currentObjID}, bson.M{
			"$pull": bson.M{"following": targetUser.ID},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow"})
			return
		}

		_, err = collection.UpdateOne(ctx, bson.M{"_id": targetUser.ID}, bson.M{
			"$pull": bson.M{"followers": currentObjID},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update followers"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Unfollowed successfully"})
	} else {
		// Follow
		_, err = collection.UpdateOne(ctx, bson.M{"_id": currentObjID}, bson.M{
			"$addToSet": bson.M{"following": targetUser.ID},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow"})
			return
		}

		_, err = collection.UpdateOne(ctx, bson.M{"_id": targetUser.ID}, bson.M{
			"$addToSet": bson.M{"followers": currentObjID},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update followers"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Followed successfully"})
	}
}

func GetUserProfile(c *gin.Context) {
	username := c.Param("username")

	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": user})
}
