package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActivateRecordRecords struct {

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 激活状态
	ActiveStatus *string `json:"active_status,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o ActivateRecordRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActivateRecordRecords struct{}"
	}

	return strings.Join([]string{"ActivateRecordRecords", string(data)}, " ")
}
