package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationsRequest Request Object
type UpdateInstanceConfigurationsRequest struct {

	// **参数解释：** 实例ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceConfigurationsRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationsRequest", string(data)}, " ")
}
