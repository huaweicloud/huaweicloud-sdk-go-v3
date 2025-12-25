package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EpResourceStatistics 企业项目关联资源的状态
type EpResourceStatistics struct {

	// **参数解释** 企业项目ID。 **取值范围** 长度[0,64]个字符.
	ExtendRelationId *string `json:"extend_relation_id,omitempty"`

	// **参数解释** 告警中的资源数。 **取值范围** 0-9999999
	Unhealthy *int32 `json:"unhealthy,omitempty"`

	// **参数解释** 资源总数。 **取值范围** 0-9999999
	Total *int32 `json:"total,omitempty"`

	// **参数解释** 已触发的资源数。 **取值范围** 0-9999999
	EventUnhealthy *int32 `json:"event_unhealthy,omitempty"`

	// **参数解释** 资源类型数。 **取值范围** 0-9999999
	Namespaces *int32 `json:"namespaces,omitempty"`
}

func (o EpResourceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpResourceStatistics struct{}"
	}

	return strings.Join([]string{"EpResourceStatistics", string(data)}, " ")
}
