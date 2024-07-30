package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns a new slice containing the records for which the predicate
// function returns true.
//
// The Filter function takes in a slice of records 'in' and a predicate
// function that takes a single record as input and returns a boolean value.
// It iterates over each record in the input slice and applies the predicate
// function to determine if the record should be included in the filtered
// list. The filtered list is then returned as a new slice.
func Filter(in []Record, predicate func(Record) bool) []Record {
	filteredList := make([]Record, 0, len(in))
	for _, record := range in {
		if predicate(record) {
			filteredList = append(filteredList, record)
		}
	}
	return filteredList
}

// ByDaysPeriod returns a predicate function that checks if the day of the
// record falls within the specified period.
//
// Parameters: - p: The DaysPeriod struct that defines the period of days.
//
// Returns: A function that takes a Record as an argument and returns true if
// the day of the record is within the period, false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return record.Day >= p.From && record.Day <= p.To
	}
}

// ByCategory is a higher-order function that returns a predicate function.
// The returned predicate function checks if the category of a record matches
// the provided category.
//
// Args:
//   c (string): The category to compare against.
//
// Returns:
//   func(Record) bool: The predicate function.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == c
	}
}

// TotalByPeriod returns the total amount of expenses for records
// that fall within the given period.
//
// Parameters:
// - in: a slice of records
// - p: the period to filter the records by
//
// Returns:
// - total: the total amount of expenses
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	for _, record := range Filter(in, ByDaysPeriod(p)) {
		total += record.Amount
	}
	return total
}

// CategoryExpenses returns the total amount of expenses for records
// in the given category and period.
// If there are no records in the list that belong to the given category,
// an error is returned.
func CategoryExpenses(records []Record, period DaysPeriod, category string) (float64, error) {
	// Filter records by category
	filteredList := Filter(records, ByCategory(category))

	// Check if there are no records in the filtered list
	if len(filteredList) == 0 {
		return 0, fmt.Errorf("unknown category %s", category)
	}

	// Calculate the total amount of expenses
	total := 0.0
	for _, record := range filteredList {
		// Check if the record falls within the given period
		if record.Day >= period.From && record.Day <= period.To {
			total += record.Amount
		}
	}

	return total, nil
}
