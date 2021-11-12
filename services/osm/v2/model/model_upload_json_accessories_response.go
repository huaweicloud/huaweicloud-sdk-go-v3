package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadJsonAccessoriesResponse struct {
	// 附件id

	AccessoryId    *string `json:"accessory_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadJsonAccessoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadJsonAccessoriesResponse struct{}"
	}

	return strings.Join([]string{"UploadJsonAccessoriesResponse", string(data)}, " ")
}
