package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAssetCategoryRequest struct {
	Body *CreateCategoryReq `json:"body,omitempty"`
}

func (o CreateAssetCategoryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetCategoryRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetCategoryRequest", string(data)}, " ")
}
