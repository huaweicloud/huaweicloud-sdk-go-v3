package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditDatabaseNewRequest Request Object
type DeleteAuditDatabaseNewRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**： 数据库ID。可在查询数据库列表接口ID字段获取。 **约束限制**： 不涉及 **取值范围**： 以查询数据库列表接口值为准，字符长度16-64。 **默认取值**： 不涉及
	Id string `json:"id"`
}

func (o DeleteAuditDatabaseNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditDatabaseNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuditDatabaseNewRequest", string(data)}, " ")
}
