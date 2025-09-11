package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmUpgradeAuditRequest Request Object
type ConfirmUpgradeAuditRequest struct {

	// **参数解释**：  资源ID。可在查询实例列表接口的resource_id字段获取。 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口获取值为准，字符长度32-64。 **默认取值**： 不涉及
	ResourceId string `json:"resource_id"`
}

func (o ConfirmUpgradeAuditRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmUpgradeAuditRequest struct{}"
	}

	return strings.Join([]string{"ConfirmUpgradeAuditRequest", string(data)}, " ")
}
