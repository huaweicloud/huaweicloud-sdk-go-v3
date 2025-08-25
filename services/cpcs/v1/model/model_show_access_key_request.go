package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessKeyRequest Request Object
type ShowAccessKeyRequest struct {

	// 访问密钥ID
	AccessKeyId string `json:"access_key_id"`

	// 应用ID
	AppId string `json:"app_id"`
}

func (o ShowAccessKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowAccessKeyRequest", string(data)}, " ")
}
