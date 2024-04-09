package models

import (
	"time"

	"example.com/rest-api/db"
)

type GSSWishlistItem struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	SpaceID   int64     `json:"spaceId"`
	StartDate string    `json:"startDate" binding:"required"`
	EndDate   string    `json:"endDate" binding:"required"`
	IsAdd     bool      `json:"isAdd"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

func GetALLWishlist() ([]GSSWishlistItem, error) {

	query := "SELECT * FROM wishlistGssTBL"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gssWishlistItems []GSSWishlistItem

	for rows.Next() {
		var wishlist GSSWishlistItem
		var isAddByte []byte
		err := rows.Scan(&wishlist.ID, &wishlist.UserID, &wishlist.SpaceID, &wishlist.StartDate, &wishlist.EndDate, &isAddByte, &wishlist.CreatedOn, &wishlist.UpdatedOn)

		if err != nil {
			return nil, err
		}

		wishlist.IsAdd = isAddByte[0] != 0 // Convert non-zero byte to true

		gssWishlistItems = append(gssWishlistItems, wishlist)
	}

	return gssWishlistItems, nil
}

func (w *GSSWishlistItem) AddToWishlist(spaceID int64) error {

	query := `
	INSERT INTO wishlistGssTBL(user_Id, space_Id, start_Date, end_Date, is_add, created_On, updated_On) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Get current time
	currentTime := time.Now()

	// Execute the SQL query with the provided parameters
	result, err := stmt.Exec(w.UserID, spaceID, w.StartDate, w.EndDate, true, currentTime, currentTime)
	if err != nil {
		return err
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Set the ID and timestamps in the wishlist item object
	w.ID = int64(id)
	w.SpaceID = int64(spaceID)
	w.CreatedOn = currentTime
	w.UpdatedOn = currentTime

	return nil
}
