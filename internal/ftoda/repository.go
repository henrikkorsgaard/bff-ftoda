package ftoda

type FTODARepository interface {
	GetSag(id int) (Sag, error)
}

/**
- there need to be a proxy that implements the interface as well

*/

// Implementation 1: Get something from the API
// Implementation 2: Get something from the database
// Implementation 3: Get something from API, if not in databaser
