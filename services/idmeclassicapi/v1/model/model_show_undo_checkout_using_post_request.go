package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUndoCheckoutUsingPostRequest Request Object
type ShowUndoCheckoutUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o ShowUndoCheckoutUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUndoCheckoutUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowUndoCheckoutUsingPostRequest", string(data)}, " ")
}
