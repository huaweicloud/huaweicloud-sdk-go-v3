package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoEnlargePolicyRequest Request Object
type ShowAutoEnlargePolicyRequest struct {

	// **参数解释：** 实例ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`
}

func (o ShowAutoEnlargePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoEnlargePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoEnlargePolicyRequest", string(data)}, " ")
}
