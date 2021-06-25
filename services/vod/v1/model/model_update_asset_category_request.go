package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAssetCategoryRequest struct {
	Body *UpdateCategoryReq `json:"body,omitempty"`
}

func (o UpdateAssetCategoryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAssetCategoryRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetCategoryRequest", string(data)}, " ")
}
