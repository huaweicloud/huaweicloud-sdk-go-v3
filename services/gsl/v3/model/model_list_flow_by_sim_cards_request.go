package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListFlowBySimCardsRequest struct {
	Body *ListFlowBySimCardsReq `json:"body,omitempty"`
}

func (o ListFlowBySimCardsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsRequest struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsRequest", string(data)}, " ")
}
