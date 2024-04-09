package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getWishlists(context *gin.Context) {
	events, err := models.GetALLWishlist()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch wishlist. Please Try again later." + err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createWishlist(context *gin.Context) {
	spaceID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var wishlist models.GSSWishlistItem
	err = context.ShouldBindJSON(&wishlist)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data." + err.Error()})
		return
	}

	userId := context.GetInt64("userId")

	wishlist.UserID = userId

	err = wishlist.AddToWishlist(spaceID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create wishlist. Try again later." + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Wishlist Created!", "event": wishlist})
}
