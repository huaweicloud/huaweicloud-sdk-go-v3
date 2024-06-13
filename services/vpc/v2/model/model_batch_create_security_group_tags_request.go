package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSecurityGroupTagsRequest Request Object
type BatchCreateSecurityGroupTagsRequest struct {

	// 安全组资源ID
	SecurityGroupId string `json:"security_group_id"`

	Body *BatchCreateSecurityGroupTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSecurityGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSecurityGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSecurityGroupTagsRequest", string(data)}, " ")
}
