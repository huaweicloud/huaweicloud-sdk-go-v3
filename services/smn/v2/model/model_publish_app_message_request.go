package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAppMessageRequest Request Object
type PublishAppMessageRequest struct {

	// Endpoint的唯一资源标识，可通过[查询Application的Endpoint列表](smn_api_58004.xml)获取该标识
	EndpointUrn string `json:"endpoint_urn"`

	Body *PublishAppMessageRequestBody `json:"body,omitempty"`
}

func (o PublishAppMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppMessageRequest struct{}"
	}

	return strings.Join([]string{"PublishAppMessageRequest", string(data)}, " ")
}
