package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditAgentRequest Request Object
type ListAuditAgentRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// 数据库ID,可在查询数据库列表接口ID字段获取。
	DbId string `json:"db_id"`

	// **参数解释**： 分页偏移量，从第一条数据偏移offset条数据后开始查询 **约束限制**： 仅支持大于等于0的整数 **取值范围**： 大于等于0 **默认取值**： 默认值为0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 每页查询记录数。 **约束限制**： 仅支持大于0的整数 **取值范围**： 大于0小于等于10000 **默认取值**： 默认值为100
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditAgentRequest struct{}"
	}

	return strings.Join([]string{"ListAuditAgentRequest", string(data)}, " ")
}
