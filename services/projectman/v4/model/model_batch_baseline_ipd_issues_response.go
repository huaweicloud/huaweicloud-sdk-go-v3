package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBaselineIpdIssuesResponse Response Object
type BatchBaselineIpdIssuesResponse struct {

	// 批量基线工作项的结果列表。
	Result *[]BatchBaselineIssueResponseResult `json:"result,omitempty"`

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 操作失败原因。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchBaselineIpdIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBaselineIpdIssuesResponse struct{}"
	}

	return strings.Join([]string{"BatchBaselineIpdIssuesResponse", string(data)}, " ")
}
