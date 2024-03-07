package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateVersionRequest Request Object
type BatchUpdateVersionRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUpdateDto `json:"body,omitempty"`
}

func (o BatchUpdateVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateVersionRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateVersionRequest", string(data)}, " ")
}
