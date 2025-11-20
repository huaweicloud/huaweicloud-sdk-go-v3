package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePremiumInstanceRequest Request Object
type UpdatePremiumInstanceRequest struct {

	// **参数解释：** 独享引擎ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	InstanceId string `json:"instance_id"`

	Body *UpdatePremiumInstanceRequestBody `json:"body,omitempty"`
}

func (o UpdatePremiumInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdatePremiumInstanceRequest", string(data)}, " ")
}
