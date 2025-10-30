package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeExtendedWeakPasswordRequest Request Object
type ChangeExtendedWeakPasswordRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	Body *ChangeExtendedWeakPasswordRequestInfo `json:"body,omitempty"`
}

func (o ChangeExtendedWeakPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeExtendedWeakPasswordRequest struct{}"
	}

	return strings.Join([]string{"ChangeExtendedWeakPasswordRequest", string(data)}, " ")
}
