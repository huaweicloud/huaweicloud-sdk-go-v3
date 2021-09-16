package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFlowBySimCardsResponse struct {
	Body           *[]SimCardsFlowVo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListFlowBySimCardsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlowBySimCardsResponse struct{}"
	}

	return strings.Join([]string{"ListFlowBySimCardsResponse", string(data)}, " ")
}
