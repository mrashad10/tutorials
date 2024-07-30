// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident creates a new resident in the city's IT system.
// It takes the resident's name, age, and address as parameters
// and returns a pointer to a Resident struct.
func NewResident(name string, age int, address map[string]string) *Resident {
	// Create a new Resident struct with the provided name, age, and address.
	resident := &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}

	// Return a pointer to the newly created Resident struct.
	return resident
}

// HasRequiredInfo checks if a resident has provided all the required
// information.
//
// A resident is considered to have provided the required information if they
// have a non-empty name and a street specified in their address.
//
// Returns true if the resident has provided all the required information,
// otherwise returns false.
func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && r.Address["street"] != ""
}

// Delete removes a resident's information from the system.
// It sets all the fields of the resident to their zero value.
func (r *Resident) Delete() {
	*r = Resident{}
}

// Count calculates the total number of residents who have provided the
// required information.
//
// Parameters:
// - residents: A slice of pointers to Resident structs.
//
// Returns:
// - The number of residents who have provided the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, r := range residents {
		if r.HasRequiredInfo() {
			count++
		}
	}

	return count
}
