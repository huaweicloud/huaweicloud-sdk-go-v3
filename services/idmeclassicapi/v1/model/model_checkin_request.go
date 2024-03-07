package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckinRequest Request Object
type CheckinRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionCheckInDto `json:"body,omitempty"`
}

func (o CheckinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckinRequest struct{}"
	}

	return strings.Join([]string{"CheckinRequest", string(data)}, " ")
}
