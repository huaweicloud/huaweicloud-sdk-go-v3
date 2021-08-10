package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCompetitionScoreRequest struct {
	Body *CreateScoresRequestModel `json:"body,omitempty"`
}

func (o CreateCompetitionScoreRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCompetitionScoreRequest struct{}"
	}

	return strings.Join([]string{"CreateCompetitionScoreRequest", string(data)}, " ")
}
