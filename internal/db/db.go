package db

import (
	"database/sql"
	"fmt"
	"math"
	"time"

	pb "github.com/ganduumar/klausapp-softwareengineer-test-task/internal/grpc/pb"
	"github.com/ganduumar/klausapp-softwareengineer-test-task/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// DB to easily access the SQLite database.
type DB struct {
	db *sql.DB
}

// NewDB for a new SQLite database initalize.
func NewDB() (*DB, error) {
	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "./internal/db/database.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return &DB{db: db}, nil
}

// Close the database connection. Easy to re-use.
func (d *DB) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

func (d *DB) GetRatingById(ratingId int32) (float64, error) {
	// Prepare SQL statement to fetch rating
	stmt, err := d.db.Prepare("SELECT rating, rating_category_id FROM ratings WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Query the rating information from the database
	var rating float64
	var ratingCategoryID int
	err = stmt.QueryRow(ratingId).Scan(&rating, &ratingCategoryID)
	if err != nil {
		return 0, fmt.Errorf("failed to query database: %v", err)
	}

	// Calculate the weighted score using CalculateRatingWithWeight
	weightedScore, err := d.CalculateRatingWithWeight(rating, ratingCategoryID)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate weighted score: %v", err)
	}

	return weightedScore, nil
}

func (d *DB) CalculateRatingWithWeight(ratingValue float64, categoryID int) (float64, error) {
	// Fetch category weight from the database
	var weight float64
	err := d.db.QueryRow("SELECT weight FROM rating_categories WHERE id = ?", categoryID).Scan(&weight)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch category weight: %v", err)
	}

	// Calculate the weighted score
	// Could but the multiplier 20 to const, but shouldn't change at current logic
	weightedScore := ratingValue * weight * 20

	return weightedScore, nil
}

func (d *DB) GetRatingsBetweenDates(startDate, endDate time.Time) (map[int32]*pb.RatingResult, error) {
	// Prepare SQL statement to fetch ratings between the specified dates
	stmt, err := d.db.Prepare(`
        SELECT 
			rc.id, 
			rc.name, 
			AVG(r.rating)
        FROM ratings r
        JOIN rating_categories rc ON r.rating_category_id = rc.id
        WHERE r.created_at BETWEEN ? AND ?
        GROUP BY rc.id, rc.name
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Run query with parameters
	rows, err := stmt.Query(startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	defer rows.Close()

	results := make(map[int32]*pb.RatingResult)

	// Iterate rows
	for rows.Next() {
		var ratingCategoryID int32
		var categoryName string
		var averageRating float64
		if err := rows.Scan(&ratingCategoryID, &categoryName, &averageRating); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Calculate the weighted score, as it's not always 1:1 ratio
		weightedScore, err := d.CalculateRatingWithWeight(averageRating, int(ratingCategoryID))
		if err != nil {
			return nil, fmt.Errorf("failed to calculate weighted score: %v", err)
		}
		// We could technically round in the weight calculation, but never know when need more accurate data
		roundedScore := math.Round(weightedScore)

		ratingResult := &pb.RatingResult{
			Id:     ratingCategoryID,
			Name:   categoryName,
			Rating: roundedScore,
		}

		results[ratingCategoryID] = ratingResult
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return results, nil
}

func (d *DB) GetRatingsBetweenDatesAggregated(startDate, endDate time.Time) (map[int32]map[string]*pb.RatingResult, error) {
	// Determine the duration between startDate and endDate
	duration := endDate.Sub(startDate)

	// Define the date format for the SQL query
	dateFormat := "2006-01-02"

	// Prepare SQL statement to fetch ratings between the specified dates
	var query string
	if duration <= 30*24*time.Hour {
		// If duration is less than or equal to 30 days
		query = `
            SELECT 
				rc.id, 
				rc.name, 
				DATE(r.created_at) AS date, 
				AVG(r.rating)
            FROM ratings r
            JOIN rating_categories rc ON r.rating_category_id = rc.id
            WHERE r.created_at BETWEEN ? AND ?
            GROUP BY rc.id, rc.name, date
        `
	} else {
		// If duration is more than 30 days
		query = `
		SELECT 
			rc.id, 
			rc.name, 
			strftime('%Y-%W', r.created_at) AS week, 
			AVG(r.rating)
		FROM ratings r
		JOIN rating_categories rc ON r.rating_category_id = rc.id
		WHERE r.created_at BETWEEN ? AND ?
		GROUP BY rc.id, rc.name, week
        `
	}

	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Query the ratings between the specified dates
	rows, err := stmt.Query(startDate.Format(dateFormat), endDate.Format(dateFormat))
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	defer rows.Close()

	results := make(map[int32]map[string]*pb.RatingResult)

	// Iterate over the rows and calculate the weighted average for each rating category
	for rows.Next() {
		var ratingCategoryID int32
		var categoryName string
		var dateOrWeek string
		var averageRating float64
		if err := rows.Scan(&ratingCategoryID, &categoryName, &dateOrWeek, &averageRating); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Calculate the weighted score using the CalculateRatingWithWeight method
		weightedScore, err := d.CalculateRatingWithWeight(averageRating, int(ratingCategoryID))
		if err != nil {
			return nil, fmt.Errorf("failed to calculate weighted score: %v", err)
		}

		// Round the weighted score to the nearest full number
		roundedScore := math.Round(weightedScore)

		// Create a RatingResult message
		ratingResult := &pb.RatingResult{
			Id:     ratingCategoryID,
			Name:   categoryName,
			Rating: roundedScore,
		}

		// Initialize the map if not already initialized
		if results[ratingCategoryID] == nil {
			results[ratingCategoryID] = make(map[string]*pb.RatingResult)
		}

		// Store the result in the map
		results[ratingCategoryID][dateOrWeek] = ratingResult
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return results, nil
}

func (d *DB) GetCategoryRatingsBetweenDatesAggregated(startDate, endDate time.Time) ([]*models.CategoryRating, error) {
	// Prepare SQL statement to fetch ratings between the specified dates
	stmt, err := d.db.Prepare(`
        SELECT r.ticket_id, c.name, r.rating_category_id, r.rating, c.weight
        FROM ratings r
        JOIN rating_categories c ON r.rating_category_id = c.id
        WHERE r.created_at BETWEEN ? AND ?
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the prepared statement
	rows, err := stmt.Query(startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}
	defer rows.Close()

	// Initialize a map to store the aggregated category ratings
	categoryRatingsMap := make(map[int32]*models.CategoryRating)

	// Iterate over the rows and calculate weighted ratings for each category of each ticket
	for rows.Next() {
		var ticketID, categoryID int32
		var categoryName string
		var rating, weight float64

		// Scan the row values into variables
		if err := rows.Scan(&ticketID, &categoryName, &categoryID, &rating, &weight); err != nil {
			return nil, fmt.Errorf("failed to scan row values: %v", err)
		}

		// Check if the ticket ID already exists in the map
		if _, ok := categoryRatingsMap[ticketID]; !ok {
			// If not, create a new entry in the map
			categoryRatingsMap[ticketID] = &models.CategoryRating{
				TicketID:      ticketID,
				Ratings:       make(map[int32]float64),
				CategoryNames: make(map[int32]string),
			}
		}

		// Store the category name in the CategoryNames map
		categoryRatingsMap[ticketID].CategoryNames[categoryID] = categoryName

		// Calculate the weighted rating for the current category of the ticket
		weightedRating := rating * weight * 20

		// Check if the category already has ratings
		if _, ok := categoryRatingsMap[ticketID].Ratings[categoryID]; !ok {
			// If not, initialize the sum and count for the category
			categoryRatingsMap[ticketID].Ratings[categoryID] = weightedRating
		} else {
			// If yes, update the sum and count for the category
			categoryRatingsMap[ticketID].Ratings[categoryID] += weightedRating
		}
	}

	// Convert the map to a slice of CategoryRating
	var categoryRatings []*models.CategoryRating
	for _, cr := range categoryRatingsMap {
		categoryRatings = append(categoryRatings, cr)
	}

	// Return the slice of aggregated category ratings
	return categoryRatings, nil
}
