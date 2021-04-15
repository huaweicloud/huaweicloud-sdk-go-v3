package model

import (
	"encoding/json"

	"strings"
)

type ImageDetectionResultDetailPolitics struct {
	Confidence *float32 `json:"confidence,omitempty"`

	Label *string `json:"label,omitempty"`

	FaceDetail *ImageDetectionResultDetailFaceDetail `json:"face_detail,omitempty"`
}

func (o ImageDetectionResultDetailPolitics) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetailPolitics struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetailPolitics", string(data)}, " ")
}
