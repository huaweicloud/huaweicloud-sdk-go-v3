package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAssetCategoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetCategoryResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAssetCategoryResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetCategoryResponse", string(data)}, " ")
}
