package sportmonks_client

type (
	// CountriesResponse struct
	CountriesResponse struct {
		Data []Country `json:"data"`
		Meta Meta      `json:"meta"`
	}

	FixtureResponse struct {
		Data Fixture `json:"data"`
	}

	// LeaguesResponse
	LeaguesResponse struct {
		Data []League `json:"data"`
		Meta Meta     `json:"meta"`
	}

	// Meta struct
	Meta struct {
		Pagination struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
		}
	}

	// PlayerResponse struct
	PlayerResponse struct {
		Data Player `json:"data"`
	}

	// SeasonsResponse
	SeasonsResponse struct {
		Data []Season `json:"data"`
		Meta Meta     `json:"meta"`
	}

	// SeasonResponse
	SeasonResponse struct {
		Data Season `json:"data"`
	}

	//SquadResponse
	SquadResponse struct {
		Data []SquadPlayer `json:"data"`
	}

	// TeamsResponse
	TeamsResponse struct {
		Data []Team `json:"data"`
	}

	// VenueResponse
	VenuesResponse struct {
		Data []Venue `json:"data"`
	}
)
