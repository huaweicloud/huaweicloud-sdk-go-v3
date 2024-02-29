package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckoutUsingPostRequest Request Object
type ShowCheckoutUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionCheckOutDto `json:"body,omitempty"`
}

func (o ShowCheckoutUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckoutUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckoutUsingPostRequest", string(data)}, " ")
}
