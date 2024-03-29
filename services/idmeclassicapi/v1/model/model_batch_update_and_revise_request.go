package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAndReviseRequest Request Object
type BatchUpdateAndReviseRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionReviseAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o BatchUpdateAndReviseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAndReviseRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAndReviseRequest", string(data)}, " ")
}
