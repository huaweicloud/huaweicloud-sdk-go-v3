package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ApplyJudgementRequest struct {
	Body *JudgementTaskRequestBody `json:"body,omitempty"`
}

func (o ApplyJudgementRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApplyJudgementRequest struct{}"
	}

	return strings.Join([]string{"ApplyJudgementRequest", string(data)}, " ")
}
