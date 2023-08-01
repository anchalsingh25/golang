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

type ByDuration []*Movie

func (a ByDuration) Len() int           { return len(a) }
func (a ByDuration) Less(i, j int) bool { return a[i].Duration > a[j].Duration }
func (a ByDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByYear []*Movie

func (a ByYear) Len() int           { return len(a) }
func (a ByYear) Less(i, j int) bool { return a[i].Year > a[j].Year }
func (a ByYear) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByRating []*Movie

func (a ByRating) Len() int           { return len(a) }
func (a ByRating) Less(i, j int) bool { return a[i].AverageRating > a[j].AverageRating }
func (a ByRating) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByNumRatings []*Movie

func (a ByNumRatings) Len() int           { return len(a) }
func (a ByNumRatings) Less(i, j int) bool { return a[i].NumberOfRatings > a[j].NumberOfRatings }
func (a ByNumRatings) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

	// Sort and print the top 5 movies based on different criteria
	sort.Sort(ByDuration(movies))
	fmt.Println("Top 5 Movies based on Duration:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Duration In Minutes: %d\n", movies[i].Title, movies[i].Duration)
	}

	sort.Sort(ByYear(movies))
	fmt.Println("\nTop 5 Movies based on Year of Release:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Year of Release: %d\n", movies[i].Title, movies[i].Year)
	}

	sort.Sort(ByRating(movies))
	fmt.Println("\nTop 5 Movies based on Average Rating:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Average Rating: %.3f\n", movies[i].Title, movies[i].AverageRating)
	}

	sort.Sort(ByNumRatings(movies))
	fmt.Println("\nTop 5 Movies based on Number of Ratings Given:")
	for i := 0; i < 5 && i < len(movies); i++ {
		fmt.Printf("%s; Number of Ratings: %d\n", movies[i].Title, movies[i].NumberOfRatings)
	}
}
