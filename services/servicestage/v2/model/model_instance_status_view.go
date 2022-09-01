package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 状态详情。
type InstanceStatusView struct {
	Status *InstanceStatusType `json:"status,omitempty" xml:"status"`

	// 正常实例副本数。
	AvailableReplica *int32 `json:"available_replica,omitempty" xml:"available_replica"`

	// 实例副本数。
	Replica *int32 `json:"replica,omitempty" xml:"replica"`

	FailDetail *InstanceFailDetail `json:"fail_detail,omitempty" xml:"fail_detail"`

	// 最近Job ID。
	LastJobId *string `json:"last_job_id,omitempty" xml:"last_job_id"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o InstanceStatusView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatusView struct{}"
	}

	return strings.Join([]string{"InstanceStatusView", string(data)}, " ")
}
