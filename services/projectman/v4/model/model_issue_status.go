package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueStatus struct {

	// 状态uuid
	Id *string `json:"id,omitempty"`

	// 状态数字id
	StatusId *int32 `json:"status_id,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`

	// 关联的工作项类型列表
	TrackerIds *[]int32 `json:"tracker_ids,omitempty"`

	StatusAttribute *StatusAttribute `json:"status_attribute,omitempty"`
}

func (o IssueStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueStatus struct{}"
	}

	return strings.Join([]string{"IssueStatus", string(data)}, " ")
}
