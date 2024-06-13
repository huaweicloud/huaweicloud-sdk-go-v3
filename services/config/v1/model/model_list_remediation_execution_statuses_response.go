package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemediationExecutionStatusesResponse Response Object
type ListRemediationExecutionStatusesResponse struct {

	// 合规规则修正执行结果查询列表。
	Value *[]RemediationExecution `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRemediationExecutionStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemediationExecutionStatusesResponse struct{}"
	}

	return strings.Join([]string{"ListRemediationExecutionStatusesResponse", string(data)}, " ")
}
