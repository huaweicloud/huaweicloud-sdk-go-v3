package model

import (
	"encoding/json"

	"strings"
)

//
type RelatedIntention struct {
	// 意图名称。

	Intention string `json:"intention"`
	// 意图置信度。

	Confidence *float64 `json:"confidence,omitempty"`
}

func (o RelatedIntention) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RelatedIntention struct{}"
	}

	return strings.Join([]string{"RelatedIntention", string(data)}, " ")
}
