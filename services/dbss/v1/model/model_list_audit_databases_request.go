package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditDatabasesRequest Request Object
type ListAuditDatabasesRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// 实例状态 - ON :开启 - OFF : 关闭
	Status *string `json:"status,omitempty"`

	// **参数解释**： 分页偏移量，从第一条数据偏移offset条数据后开始查询 **约束限制**： 仅支持大于等于0的整数 **取值范围**： 大于等于0 **默认取值**： 默认值为0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 每页查询记录数。 **约束限制**： 仅支持大于0的整数 **取值范围**： 大于0小于等于10000 **默认取值**： 默认值为100
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListAuditDatabasesRequest", string(data)}, " ")
}
