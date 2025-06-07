package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/osdeving/poker-timeleft-api/database"
	"github.com/osdeving/poker-timeleft-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection() *mongo.Collection {
	return database.MongoClient.Database("poker-timeleft").Collection("tournaments")
}

func CreateTournament(c *gin.Context) {
	var t models.Tournament

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getCollection()
	result, err := collection.InsertOne(ctx, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, t)
}

func GetAllTournaments(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := getCollection()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var ts []models.Tournament
	if err := cursor.All(ctx, &ts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ts)
}

func GetTournament(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getCollection()
	var t models.Tournament
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&t)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Torneio não encontrado"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func UpdateTournament(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var t models.Tournament
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getCollection()
	update := bson.M{"$set": bson.M{"name": t.Name, "note": t.Note}}
	_, err = collection.UpdateByID(ctx, objID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t.ID = objID
	c.JSON(http.StatusOK, t)
}

func DeleteTournament(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getCollection()
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
