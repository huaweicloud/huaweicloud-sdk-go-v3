package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceGroupRequest Request Object
type DeleteResourceGroupRequest struct {

	// **参数解释** 资源分组ID。 **约束限制** 不涉及。 **取值范围** 以\"rg\"开头，后面跟着22个字母或数字。 **默认取值** 不涉及。
	GroupId string `json:"group_id"`
}

func (o DeleteResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupRequest", string(data)}, " ")
}
