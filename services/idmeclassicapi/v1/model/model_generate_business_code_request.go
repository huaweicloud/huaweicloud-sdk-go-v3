package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateBusinessCodeRequest Request Object
type GenerateBusinessCodeRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`
}

func (o GenerateBusinessCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateBusinessCodeRequest struct{}"
	}

	return strings.Join([]string{"GenerateBusinessCodeRequest", string(data)}, " ")
}
