package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DecServiceFlavorItemDetail struct {

	// **参数解释**: 类型编码。 **取值范围**: 不涉及。
	TypeCode *string `json:"type_code,omitempty"`

	// **参数解释**: 规格编码。 **取值范围**: 不涉及。
	Code *string `json:"code,omitempty"`
}

func (o DecServiceFlavorItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecServiceFlavorItemDetail struct{}"
	}

	return strings.Join([]string{"DecServiceFlavorItemDetail", string(data)}, " ")
}
