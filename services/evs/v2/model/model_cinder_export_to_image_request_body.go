package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CinderExportToImageRequestBody struct {
	OsVolumeUploadImage *CinderExportToImageOption `json:"os-volume_upload_image"`
}

func (o CinderExportToImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderExportToImageRequestBody struct{}"
	}

	return strings.Join([]string{"CinderExportToImageRequestBody", string(data)}, " ")
}
