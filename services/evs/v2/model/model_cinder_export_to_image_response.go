package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CinderExportToImageResponse struct {
	OsVolumeUploadImage *Image `json:"os-volume_upload_image,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o CinderExportToImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderExportToImageResponse struct{}"
	}

	return strings.Join([]string{"CinderExportToImageResponse", string(data)}, " ")
}
