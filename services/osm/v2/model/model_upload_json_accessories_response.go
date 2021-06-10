package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UploadJsonAccessoriesResponse struct {
	// 附件id

	AccessoryId    *string `json:"accessory_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadJsonAccessoriesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadJsonAccessoriesResponse struct{}"
	}

	return strings.Join([]string{"UploadJsonAccessoriesResponse", string(data)}, " ")
}
