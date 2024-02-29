package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchCheckoutUsingPostRequest Request Object
type ShowBatchCheckoutUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionCheckOutDto `json:"body,omitempty"`
}

func (o ShowBatchCheckoutUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCheckoutUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchCheckoutUsingPostRequest", string(data)}, " ")
}
