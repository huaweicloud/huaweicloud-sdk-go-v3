package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckoutAndUpdateUsingPostRequest Request Object
type ShowCheckoutAndUpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionCheckoutAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o ShowCheckoutAndUpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckoutAndUpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckoutAndUpdateUsingPostRequest", string(data)}, " ")
}
