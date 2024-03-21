package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStateRequest Request Object
type UpdateStateRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto `json:"body,omitempty"`
}

func (o UpdateStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStateRequest struct{}"
	}

	return strings.Join([]string{"UpdateStateRequest", string(data)}, " ")
}
