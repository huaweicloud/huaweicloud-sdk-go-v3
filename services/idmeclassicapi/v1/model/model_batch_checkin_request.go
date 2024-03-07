package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckinRequest Request Object
type BatchCheckinRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionCheckInDto `json:"body,omitempty"`
}

func (o BatchCheckinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckinRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckinRequest", string(data)}, " ")
}
