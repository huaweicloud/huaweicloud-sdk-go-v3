package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAddonInstanceRequest Request Object
type UpdateAddonInstanceRequest struct {

	// **参数解释**： 插件实例ID。 **约束限制**： 不涉及 **取值范围**： UUID格式，长度范围1~255位。 **默认取值**： 不涉及
	Id string `json:"id"`

	Body *InstanceRequest `json:"body,omitempty"`
}

func (o UpdateAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateAddonInstanceRequest", string(data)}, " ")
}
