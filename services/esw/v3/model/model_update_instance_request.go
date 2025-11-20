package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRequest Request Object
type UpdateInstanceRequest struct {

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
