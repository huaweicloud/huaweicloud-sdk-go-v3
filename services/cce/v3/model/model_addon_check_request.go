package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonCheckRequest 插件检查-request结构体
type AddonCheckRequest struct {

	// **参数解释：** API类型 **约束限制：** 该值不可修改 **取值范围：** 固定值\"AddonCheck\" **默认取值：** AddonCheck
	Kind string `json:"kind"`

	// **参数解释：** API版本 **约束限制：** 该值不可修改 **取值范围：** 固定值\"v3\" **默认取值：** v3
	ApiVersion string `json:"apiVersion"`

	Spec *AddonCheckSpec `json:"spec"`
}

func (o AddonCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonCheckRequest struct{}"
	}

	return strings.Join([]string{"AddonCheckRequest", string(data)}, " ")
}
