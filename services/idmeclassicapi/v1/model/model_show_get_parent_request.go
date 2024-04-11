package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetParentRequest Request Object
type ShowGetParentRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoQueryParentDto `json:"body,omitempty"`
}

func (o ShowGetParentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetParentRequest struct{}"
	}

	return strings.Join([]string{"ShowGetParentRequest", string(data)}, " ")
}
