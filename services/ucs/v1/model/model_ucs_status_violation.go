package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsStatusViolation struct {

	// 进行审计的集群ID
	ClusterID *string `json:"clusterID,omitempty"`

	// 审计时间
	AuditTimestamp *string `json:"auditTimestamp,omitempty"`

	// 违规状态列表
	ClusterViolations *[]StatusViolation `json:"clusterViolations,omitempty"`

	// 拦截事件列表
	ClusterEvents *[]StatusEvent `json:"clusterEvents,omitempty"`

	// 告警事件列表
	WarnEvents *[]StatusEvent `json:"warnEvents,omitempty"`
}

func (o UcsStatusViolation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsStatusViolation struct{}"
	}

	return strings.Join([]string{"UcsStatusViolation", string(data)}, " ")
}
