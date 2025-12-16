package pokeapi

// a struct that models the API response
type LocationAreasResp struct{
	Count	 int		`json:"count"`
	Next	 *string 	`json:"next"`
	Previous *string 	`json:"previous"`
	Results []struct{
		Name	string `json:"name"`
		URL		string `json:"url"`
	}`json:"results"`
}
