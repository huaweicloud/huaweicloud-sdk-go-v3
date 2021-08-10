package model

import (
	"encoding/json"

	"strings"
)

type CompareFace struct {
	BoundingBox *BoundingBox `json:"bounding_box"`
}

func (o CompareFace) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompareFace struct{}"
	}

	return strings.Join([]string{"CompareFace", string(data)}, " ")
}
