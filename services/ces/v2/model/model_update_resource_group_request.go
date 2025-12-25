package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceGroupRequest Request Object
type UpdateResourceGroupRequest struct {

	// **参数解释** 资源分组ID。 **约束限制** 不涉及。 **取值范围** 以\"rg\"开头，后跟22位由字母或数字组成的字符串。 **默认取值** 不涉及。
	GroupId string `json:"group_id"`

	Body *PutResourceGroupReq `json:"body,omitempty"`
}

func (o UpdateResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupRequest", string(data)}, " ")
}
