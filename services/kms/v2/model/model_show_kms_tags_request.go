package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowKmsTagsRequest struct {

	// 密钥ID
	KeyId string `json:"key_id" xml:"key_id"`
}

func (o ShowKmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKmsTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKmsTagsRequest", string(data)}, " ")
}
