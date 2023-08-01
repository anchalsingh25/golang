package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Movie struct {
	Title           string
	Duration        int
	Year            int
	AverageRating   float64
	NumberOfRatings int
}

type ByDurationAndTitle []*Movie

func (a ByDurationAndTitle) Len() int      { return len(a) }
func (a ByDurationAndTitle) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDurationAndTitle) Less(i, j int) bool {
	if a[i].Duration == a[j].Duration {
		return a[i].Title < a[j].Title
	}
	return a[i].Duration > a[j].Duration
}

type ByYearAndTitle []*Movie

func (a ByYearAndTitle) Len() int      { return len(a) }
func (a ByYearAndTitle) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByYearAndTitle) Less(i, j int) bool {
	if a[i].Year == a[j].Year {
		return a[i].Title < a[j].Title
	}
	return a[i].Year > a[j].Year
}

type ByRatingAndTitle []*Movie

func (a ByRatingAndTitle) Len() int      { return len(a) }
func (a ByRatingAndTitle) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRatingAndTitle) Less(i, j int) bool {
	if a[i].AverageRating == a[j].AverageRating {
		return a[i].Title < a[j].Title
	}
	return a[i].AverageRating > a[j].AverageRating
}

type ByNumRatingsAndTitle []*Movie

func (a ByNumRatingsAndTitle) Len() int      { return len(a) }
func (a ByNumRatingsAndTitle) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByNumRatingsAndTitle) Less(i, j int) bool {
	if a[i].NumberOfRatings == a[j].NumberOfRatings {
		return a[i].Title < a[j].Title
	}
	return a[i].NumberOfRatings > a[j].NumberOfRatings
}

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error Opening File", err)
		os.Exit(1)
	}
	return file
}

func giveFileData(file *os.File) [][]string {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading records", err)
		os.Exit(1)
	}
	return records
}

func main() {
	moviesFile := openFile("movies.csv")
	defer moviesFile.Close()

	ratingsFile := openFile("ratings.csv")
	defer ratingsFile.Close()

	moviesData := giveFileData(moviesFile)
	ratingsData := giveFileData(ratingsFile)

	movies := make([]*Movie, len(moviesData)-1)

	for i, record := range moviesData[1:] {
		duration, _ := strconv.Atoi(record[6])
		year, _ := strconv.Atoi(record[2])

		movies[i] = &Movie{
			Title:           record[1],
			Duration:        duration,
			Year:            year,
			AverageRating:   0.0,
			NumberOfRatings: 0,
		}
	}

	for _, record := range ratingsData[1:] {
		movieID, _ := strconv.Atoi(record[0])
		averageRating, _ := strconv.ParseFloat(record[1], 64)
		numRatings, _ := strconv.Atoi(record[2])

		movies[movieID-1].AverageRating = averageRating
		movies[movieID-1].NumberOfRatings = numRatings
	}

	sort.Sort(ByDurationAndTitle(movies))
	fmt.Println("Top 5 Movies based on Duration:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Duration In Minutes: %d\n", movies[i].Title, movies[i].Duration)
	}

	sort.Sort(ByYearAndTitle(movies))
	fmt.Println("\nTop 5 Movies based on Year of Release:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Year of Release: %d\n", movies[i].Title, movies[i].Year)
	}

	sort.Sort(ByRatingAndTitle(movies))
	fmt.Println("\nTop 5 Movies based on Average Rating:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Average Rating: %.3f\n", movies[i].Title, movies[i].AverageRating)
	}

	sort.Sort(ByNumRatingsAndTitle(movies))
	fmt.Println("\nTop 5 Movies based on Number of Ratings Given:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Number of Ratings: %d\n", movies[i].Title, movies[i].NumberOfRatings)
	}
}
