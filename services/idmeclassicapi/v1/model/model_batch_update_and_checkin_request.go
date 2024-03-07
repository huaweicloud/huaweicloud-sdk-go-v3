package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAndCheckinRequest Request Object
type BatchUpdateAndCheckinRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUpdateAndCheckinDtoVersionModel `json:"body,omitempty"`
}

func (o BatchUpdateAndCheckinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAndCheckinRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAndCheckinRequest", string(data)}, " ")
}
