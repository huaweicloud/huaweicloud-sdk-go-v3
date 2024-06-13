package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityGroupTagsRequest Request Object
type BatchDeleteSecurityGroupTagsRequest struct {

	// 安全组资源ID
	SecurityGroupId string `json:"security_group_id"`

	Body *BatchDeleteSecurityGroupTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSecurityGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityGroupTagsRequest", string(data)}, " ")
}
