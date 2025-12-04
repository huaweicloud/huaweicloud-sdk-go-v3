package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKmsKeyRequest Request Object
type DeleteKmsKeyRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 当前已绑定密钥的资源ID。
	Id string `json:"id"`
}

func (o DeleteKmsKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKmsKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteKmsKeyRequest", string(data)}, " ")
}
