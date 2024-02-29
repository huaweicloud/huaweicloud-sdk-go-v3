package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckinUsingPostRequest Request Object
type ShowCheckinUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionCheckInDto `json:"body,omitempty"`
}

func (o ShowCheckinUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckinUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckinUsingPostRequest", string(data)}, " ")
}
