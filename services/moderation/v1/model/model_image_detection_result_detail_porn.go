package model

import (
	"encoding/json"

	"strings"
)

type ImageDetectionResultDetailPorn struct {
	Confidence *float32 `json:"confidence,omitempty"`

	Label *string `json:"label,omitempty"`
}

func (o ImageDetectionResultDetailPorn) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetailPorn struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetailPorn", string(data)}, " ")
}
