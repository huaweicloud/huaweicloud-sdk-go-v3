package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAndCheckinRequest Request Object
type UpdateAndCheckinRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionUpdateAndCheckinDtoVersionModel `json:"body,omitempty"`
}

func (o UpdateAndCheckinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAndCheckinRequest struct{}"
	}

	return strings.Join([]string{"UpdateAndCheckinRequest", string(data)}, " ")
}
