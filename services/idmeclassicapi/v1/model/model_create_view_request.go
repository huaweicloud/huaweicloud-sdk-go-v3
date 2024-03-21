package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateViewRequest Request Object
type CreateViewRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoMultiViewModelVersionViewCreateDto `json:"body,omitempty"`
}

func (o CreateViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateViewRequest struct{}"
	}

	return strings.Join([]string{"CreateViewRequest", string(data)}, " ")
}
