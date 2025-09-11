package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlDetailRequest Request Object
type ShowSqlDetailRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**：  SQL语句ID。可在SQL列表查询接口ID字段获取。 **约束限制**： 不涉及 **取值范围**： 以SQL列表查询接口获取值为准。 **默认取值**： 不涉及
	Id string `json:"id"`
}

func (o ShowSqlDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlDetailRequest", string(data)}, " ")
}
