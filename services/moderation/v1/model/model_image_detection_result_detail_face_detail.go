package model

import (
	"encoding/json"

	"strings"
)

type ImageDetectionResultDetailFaceDetail struct {
	H *int32 `json:"h,omitempty"`

	W *int32 `json:"w,omitempty"`

	X *int32 `json:"x,omitempty"`

	Y *int32 `json:"y,omitempty"`
}

func (o ImageDetectionResultDetailFaceDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetailFaceDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetailFaceDetail", string(data)}, " ")
}
