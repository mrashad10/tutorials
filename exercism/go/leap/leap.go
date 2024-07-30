package leap

// IsLeapYear determines whether the given year is a leap year.
//
// year: The year to be checked.
// bool: Returns true if the year is a leap year, false otherwise.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}
