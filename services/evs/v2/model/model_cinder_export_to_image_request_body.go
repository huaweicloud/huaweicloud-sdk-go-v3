package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CinderExportToImageRequestBody struct {
	OsVolumeUploadImage *CinderExportToImageOption `json:"os-volume_upload_image"`
}

func (o CinderExportToImageRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CinderExportToImageRequestBody struct{}"
	}

	return strings.Join([]string{"CinderExportToImageRequestBody", string(data)}, " ")
}
