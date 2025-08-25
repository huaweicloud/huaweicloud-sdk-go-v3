package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessKeyRequest Request Object
type DeleteAccessKeyRequest struct {

	// 访问密钥ID
	AccessKeyId string `json:"access_key_id"`

	// 应用ID
	AppId string `json:"app_id"`
}

func (o DeleteAccessKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAccessKeyRequest", string(data)}, " ")
}
