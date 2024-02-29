package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchCheckoutAndUpdateUsingPostRequest Request Object
type ShowBatchCheckoutAndUpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionCheckoutAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o ShowBatchCheckoutAndUpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCheckoutAndUpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchCheckoutAndUpdateUsingPostRequest", string(data)}, " ")
}
