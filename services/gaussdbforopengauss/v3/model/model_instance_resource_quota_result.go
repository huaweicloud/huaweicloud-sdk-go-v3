package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceResourceQuotaResult struct {

	// **参数解释**: 资源类型 **取值范围**: instance：实例类型资源。
	Type *string `json:"type,omitempty"`

	// **参数解释**: 资源的总配额 **取值范围**: 不涉及
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释**: 已使用的资源配额 **取值范围**: 不涉及
	Used *int32 `json:"used,omitempty"`
}

func (o InstanceResourceQuotaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceResourceQuotaResult struct{}"
	}

	return strings.Join([]string{"InstanceResourceQuotaResult", string(data)}, " ")
}
