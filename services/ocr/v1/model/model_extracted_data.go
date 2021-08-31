package model

import (
	"encoding/json"

	"strings"
)

//
type ExtractedData struct {
	MathInfo *MathInfo `json:"math_info,omitempty"`
}

func (o ExtractedData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExtractedData struct{}"
	}

	return strings.Join([]string{"ExtractedData", string(data)}, " ")
}
