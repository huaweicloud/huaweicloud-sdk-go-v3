package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAndReviseRequest Request Object
type UpdateAndReviseRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionReviseAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o UpdateAndReviseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAndReviseRequest struct{}"
	}

	return strings.Join([]string{"UpdateAndReviseRequest", string(data)}, " ")
}
