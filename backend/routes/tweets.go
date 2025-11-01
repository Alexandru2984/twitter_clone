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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTweet(c *gin.Context) {
	var req models.CreateTweetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("userID")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get username
	userCollection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	tweet := models.Tweet{
		UserID:    objID,
		Username:  user.Username,
		Content:   req.Content,
		Likes:     []primitive.ObjectID{},
		Retweets:  []primitive.ObjectID{},
		Replies:   []primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	collection := config.DB.Collection("tweets")
	result, err := collection.InsertOne(ctx, tweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
		return
	}

	tweet.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, gin.H{"tweet": tweet})
}

func GetTweets(c *gin.Context) {
	collection := config.DB.Collection("tweets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(50)
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweets"})
		return
	}
	defer cursor.Close(ctx)

	var tweets []models.Tweet
	if err = cursor.All(ctx, &tweets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode tweets"})
		return
	}

	if tweets == nil {
		tweets = []models.Tweet{}
	}

	c.JSON(http.StatusOK, gin.H{"tweets": tweets})
}

func GetUserTweets(c *gin.Context) {
	username := c.Param("username")

	collection := config.DB.Collection("tweets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := collection.Find(ctx, bson.M{"username": username}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweets"})
		return
	}
	defer cursor.Close(ctx)

	var tweets []models.Tweet
	if err = cursor.All(ctx, &tweets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode tweets"})
		return
	}

	if tweets == nil {
		tweets = []models.Tweet{}
	}

	c.JSON(http.StatusOK, gin.H{"tweets": tweets})
}

func LikeTweet(c *gin.Context) {
	tweetID := c.Param("id")
	userID := c.GetString("userID")

	tweetObjID, err := primitive.ObjectIDFromHex(tweetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tweet ID"})
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := config.DB.Collection("tweets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if already liked
	var tweet models.Tweet
	err = collection.FindOne(ctx, bson.M{"_id": tweetObjID}).Decode(&tweet)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	// Toggle like
	liked := false
	for _, likeID := range tweet.Likes {
		if likeID == userObjID {
			liked = true
			break
		}
	}

	var update bson.M
	if liked {
		update = bson.M{"$pull": bson.M{"likes": userObjID}}
	} else {
		update = bson.M{"$addToSet": bson.M{"likes": userObjID}}
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": tweetObjID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tweet like toggled"})
}

func DeleteTweet(c *gin.Context) {
	tweetID := c.Param("id")
	userID := c.GetString("userID")

	tweetObjID, err := primitive.ObjectIDFromHex(tweetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tweet ID"})
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := config.DB.Collection("tweets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if tweet belongs to user
	var tweet models.Tweet
	err = collection.FindOne(ctx, bson.M{"_id": tweetObjID}).Decode(&tweet)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	if tweet.UserID != userObjID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own tweets"})
		return
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": tweetObjID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tweet deleted"})
}
