package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKmsKeyRequest Request Object
type ShowKmsKeyRequest struct {

	// Topic的唯一的资源标识。可以通过[查看主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`
}

func (o ShowKmsKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKmsKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowKmsKeyRequest", string(data)}, " ")
}
