package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateViewRequest Request Object
type BatchCreateViewRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListMultiViewModelVersionViewCreateDto `json:"body,omitempty"`
}

func (o BatchCreateViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateViewRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateViewRequest", string(data)}, " ")
}
