// internal/grpc/server/server.go

package server

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	db "github.com/ganduumar/klausapp-softwareengineer-test-task/internal/db"
	pb "github.com/ganduumar/klausapp-softwareengineer-test-task/internal/grpc/pb"
	_ "github.com/mattn/go-sqlite3"
)

type RatingsService struct {
	pb.UnimplementedRatingServiceServer
	db *db.DB
}

func NewRatingsService(db *db.DB) *RatingsService {
	return &RatingsService{
		db: db,
	}
}

func (u *RatingsService) GetRatingByIdWithWeight(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetRatingsResponse, error) {
	// Assuming u.db is an instance of your database connection
	// and req contains the rating ID

	// Get the rating ID from the request
	ratingID := req.Id

	// Call the GetRatingByIdWithWeight method
	score, err := u.db.GetRatingById(ratingID)
	if err != nil {
		return nil, err
	}

	// Print the score (for testing purposes)
	fmt.Println("Rating score:", score)

	// Return an empty response with no error
	return &pb.GetRatingsResponse{
		Rating: []*pb.Rating{},
	}, nil
}

func (u *RatingsService) GetRatingsBetweenDates(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetRatingsResponse, error) {
	startDateStr := req.StartDate
	endDateStr := req.EndDate

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	// Call the database request method
	weightedAverages, err := u.db.GetRatingsBetweenDates(startDate, endDate)
	if err != nil {
		return nil, err
	}

	ratings := make([]*pb.Rating, 0, len(weightedAverages))
	for id, result := range weightedAverages {
		rating := &pb.Rating{
			Id:     id,
			Name:   result.Name,
			Rating: int32(result.Rating),
		}
		ratings = append(ratings, rating)
	}

	return &pb.GetRatingsResponse{
		Rating: ratings,
	}, nil
}

func (u *RatingsService) GetRatingsBetweenDatesAggregated(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetRatingsAggregatedResponse, error) {
	startDateStr := req.StartDate
	endDateStr := req.EndDate

	// Parse the start and end dates into time.Time objects
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	// Call the GetRatingsBetweenDate method
	weightedAverages, err := u.db.GetRatingsBetweenDatesAggregated(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Prepare a map to store the grouped ratings
	ratingGroups := make(map[int32]*pb.RatingGroup)

	// Iterate over the weightedAverages map and group the results by ID
	for _, resultMap := range weightedAverages {
		for dateOrWeek, result := range resultMap {
			// Check if the rating group already exists in the map
			group, exists := ratingGroups[result.Id]
			if !exists {
				// If the group doesn't exist, create a new one with defaults
				group = &pb.RatingGroup{
					Id:    result.Id,
					Name:  result.Name,
					Count: 1,
					Ratings: []*pb.RatingDetail{{
						DateOrWeek: dateOrWeek,
						Rating:     int32(result.Rating),
					}},
				}
			} else {
				// If the group exists, update the count and add the rating detail
				group.Count++
				group.Ratings = append(group.Ratings, &pb.RatingDetail{
					DateOrWeek: dateOrWeek,
					Rating:     int32(result.Rating),
				})
			}

			ratingGroups[result.Id] = group
		}
	}

	// Prepare a slice (new array) to store the rating groups in the desired order
	var ratings []*pb.RatingGroup

	// Iterate over the ratingGroups map and append the groups to the slice
	for _, group := range ratingGroups {
		ratings = append(ratings, group)
	}

	// Technically would be nice to return all responses by id first and other for json, but it keeps fighting me
	// Should want to return the ratings in the response with id, name, count, and ratings in the correct order
	return &pb.GetRatingsAggregatedResponse{
		RatingGroup: ratings,
	}, nil
}

func (u *RatingsService) GetOverallRatingsBetweenDates(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetOverallRatingsResponse, error) {
	startDateStr := req.StartDate
	endDateStr := req.EndDate

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	if endDate.Before(startDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	weightedAverages, err := u.db.GetRatingsBetweenDates(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Calculate the overall weighted average rating
	var totalWeightedRating float64
	var totalWeight float64
	for _, rating := range weightedAverages {
		// Accumulate the weighted score and total weight
		totalWeightedRating += rating.Rating
		totalWeight++
	}

	// Calculate the overall average rating
	overallRating := totalWeightedRating / totalWeight

	// Return the overall average rating in the response
	return &pb.GetOverallRatingsResponse{
		OverallRating: int32(overallRating),
	}, nil
}

func (u *RatingsService) GetOverallRatingsChangeBetweenDates(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetOverallRatingsChangeResponse, error) {
	// Extract the start and end dates from the request
	startDateStr := req.StartDate
	endDateStr := req.EndDate

	// Parse the start and end dates into time.Time objects
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	if endDate.Before(startDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	// Call the GetRatingsBetweenDate method to fetch ratings between the specified dates
	// I re-use once declared function for ease of use, technically could just take first and last.
	weightedAveragesStart, err := u.db.GetRatingsBetweenDates(startDate, startDate)
	if err != nil {
		return nil, err
	}

	// Calculate the overall weighted average rating at the start date
	var totalWeightedRatingStart float64
	var totalWeightStart float64
	for _, rating := range weightedAveragesStart {
		// Accumulate the weighted score and total weight
		totalWeightedRatingStart += rating.Rating
		totalWeightStart++
	}
	overallRatingStart := totalWeightedRatingStart / totalWeightStart

	// Database request for data
	// Could be a cause for error. I might only want the latest one in general.
	weightedAveragesEnd, err := u.db.GetRatingsBetweenDates(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Calculate the overall weighted average rating at the end date
	var totalWeightedRatingEnd float64
	var totalWeightEnd float64
	for _, rating := range weightedAveragesEnd {
		// Accumulate the weighted score and total weight
		totalWeightedRatingEnd += rating.Rating
		totalWeightEnd++
	}
	overallRatingEnd := totalWeightedRatingEnd / totalWeightEnd

	// Calculate the percentage change
	var percentageChange float64
	if overallRatingStart == 0 || math.IsNaN(overallRatingStart) {
		// If start rating is zero, return the percentage change based on the end rating. Otherwise it can become NaN
		percentageChange = overallRatingEnd
	} else {
		percentageChange = ((overallRatingEnd - overallRatingStart) / overallRatingStart)
	}

	// Convert the percentage change to int32. Small fix due to my own handling.
	if percentageChange < float64(math.MinInt32) {
		percentageChange = math.MinInt32
	} else if percentageChange > float64(math.MaxInt32) {
		percentageChange = math.MaxInt32
	}

	return &pb.GetOverallRatingsChangeResponse{
		PercentageChange: int32(percentageChange),
	}, nil
}

func (s *RatingsService) GetCategoryRatingsBetweenDatesAggregated(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetCategoryRatingsAggregatedResponse, error) {
	startDateStr := req.StartDate
	endDateStr := req.EndDate

	// Parse the start and end dates into time.Time objects
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	// Database request
	categoryRatings, err := s.db.GetCategoryRatingsBetweenDatesAggregated(startDate, endDate)
	if err != nil {
		return nil, err
	}

	var categoryRatingResponses []*pb.CategoryRating

	// There could be a nicer way to create the detail items, but this is the current result.
	for _, cr := range categoryRatings {
		var ratingCategoryDetails []*pb.RatingCategoryDetail

		for categoryID, avgRating := range cr.Ratings {
			categoryName, ok := cr.CategoryNames[categoryID]
			if !ok {
				// Just in case our SQL request was bad, we throw error. Could handle better (Ignore line).
				return nil, fmt.Errorf("category name not found for category ID: %d", categoryID)
			}

			// Create a RatingCategoryDetail message for each category
			ratingCategoryDetail := &pb.RatingCategoryDetail{
				Id:            categoryID,
				Name:          categoryName,
				AverageRating: int32(avgRating),
			}
			ratingCategoryDetails = append(ratingCategoryDetails, ratingCategoryDetail)
		}

		categoryRating := &pb.CategoryRating{
			Id:                   cr.TicketID,
			RatingCategoryDetail: ratingCategoryDetails,
		}
		categoryRatingResponses = append(categoryRatingResponses, categoryRating)
	}

	return &pb.GetCategoryRatingsAggregatedResponse{
		CategoryRating: categoryRatingResponses,
	}, nil
}
