package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateCompetitionScoreRequest struct {
	Body *UpdateScoreRequestModel `json:"body,omitempty"`
}

func (o UpdateCompetitionScoreRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCompetitionScoreRequest struct{}"
	}

	return strings.Join([]string{"UpdateCompetitionScoreRequest", string(data)}, " ")
}
