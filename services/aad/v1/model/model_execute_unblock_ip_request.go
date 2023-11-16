package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteUnblockIpRequest Request Object
type ExecuteUnblockIpRequest struct {

	// 租户id
	DomainId string `json:"domain_id"`

	Body *ExecuteUnblockIpRequestBody `json:"body,omitempty"`
}

func (o ExecuteUnblockIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUnblockIpRequest struct{}"
	}

	return strings.Join([]string{"ExecuteUnblockIpRequest", string(data)}, " ")
}
