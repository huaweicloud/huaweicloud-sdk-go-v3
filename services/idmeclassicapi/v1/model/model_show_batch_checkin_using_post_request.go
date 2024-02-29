package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchCheckinUsingPostRequest Request Object
type ShowBatchCheckinUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionCheckInDto `json:"body,omitempty"`
}

func (o ShowBatchCheckinUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCheckinUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchCheckinUsingPostRequest", string(data)}, " ")
}
