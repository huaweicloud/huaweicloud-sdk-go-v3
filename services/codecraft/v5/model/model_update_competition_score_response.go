package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCompetitionScoreResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCompetitionScoreResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCompetitionScoreResponse struct{}"
	}

	return strings.Join([]string{"UpdateCompetitionScoreResponse", string(data)}, " ")
}
