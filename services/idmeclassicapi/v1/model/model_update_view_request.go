package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateViewRequest Request Object
type UpdateViewRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionUpdateViewDtoMultiViewModelViewUpdateAttrDto `json:"body,omitempty"`
}

func (o UpdateViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateViewRequest struct{}"
	}

	return strings.Join([]string{"UpdateViewRequest", string(data)}, " ")
}
