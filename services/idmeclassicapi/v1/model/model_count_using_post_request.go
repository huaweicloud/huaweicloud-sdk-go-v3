package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountUsingPostRequest Request Object
type CountUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoQueryRequestCountVo `json:"body,omitempty"`
}

func (o CountUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountUsingPostRequest struct{}"
	}

	return strings.Join([]string{"CountUsingPostRequest", string(data)}, " ")
}
