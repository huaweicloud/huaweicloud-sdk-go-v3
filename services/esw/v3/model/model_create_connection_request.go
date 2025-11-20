package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionRequest Request Object
type CreateConnectionRequest struct {

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	InstanceId string `json:"instance_id"`

	Body *CreateConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequest", string(data)}, " ")
}
