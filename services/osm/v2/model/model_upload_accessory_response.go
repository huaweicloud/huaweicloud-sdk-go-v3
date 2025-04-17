package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAccessoryResponse Response Object
type UploadAccessoryResponse struct {

	// 附件id
	AccessoryId    *string `json:"accessory_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadAccessoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAccessoryResponse struct{}"
	}

	return strings.Join([]string{"UploadAccessoryResponse", string(data)}, " ")
}
