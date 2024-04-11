package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTargetRequest Request Object
type DeleteTargetRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoGenericLinkTypeModifierDto `json:"body,omitempty"`
}

func (o DeleteTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTargetRequest struct{}"
	}

	return strings.Join([]string{"DeleteTargetRequest", string(data)}, " ")
}
