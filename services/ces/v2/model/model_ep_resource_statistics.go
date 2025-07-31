package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EpResourceStatistics 企业项目关联资源的状态
type EpResourceStatistics struct {

	// 企业项目ID
	ExtendRelationId *string `json:"extend_relation_id,omitempty"`

	// 告警中的资源数
	Unhealthy *int32 `json:"unhealthy,omitempty"`

	// 资源总数
	Total *int32 `json:"total,omitempty"`

	// 已触发的资源数
	EventUnhealthy *int32 `json:"event_unhealthy,omitempty"`

	// 资源类型数
	Namespaces *int32 `json:"namespaces,omitempty"`
}

func (o EpResourceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpResourceStatistics struct{}"
	}

	return strings.Join([]string{"EpResourceStatistics", string(data)}, " ")
}
