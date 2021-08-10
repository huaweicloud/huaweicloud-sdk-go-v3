package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CompareFaceByUrlRequest struct {
	Body *FaceCompareUrlReq `json:"body,omitempty"`
}

func (o CompareFaceByUrlRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompareFaceByUrlRequest struct{}"
	}

	return strings.Join([]string{"CompareFaceByUrlRequest", string(data)}, " ")
}
