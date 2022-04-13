package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 状态详情。
type InstanceStatusView struct {
	Status *InstanceStatusType `json:"status,omitempty"`
	// 正常实例副本数。

	AvailableReplica *int32 `json:"available_replica,omitempty"`
	// 实例副本数。

	Replica *int32 `json:"replica,omitempty"`

	FailDetail *InstanceFailDetail `json:"fail_detail,omitempty"`
	// 最近Job ID。

	LastJobId *string `json:"last_job_id,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o InstanceStatusView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatusView struct{}"
	}

	return strings.Join([]string{"InstanceStatusView", string(data)}, " ")
}
