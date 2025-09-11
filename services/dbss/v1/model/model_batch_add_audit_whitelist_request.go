package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddAuditWhitelistRequest Request Object
type BatchAddAuditWhitelistRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	Body *AddWhitelistRequest `json:"body,omitempty"`
}

func (o BatchAddAuditWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddAuditWhitelistRequest struct{}"
	}

	return strings.Join([]string{"BatchAddAuditWhitelistRequest", string(data)}, " ")
}
