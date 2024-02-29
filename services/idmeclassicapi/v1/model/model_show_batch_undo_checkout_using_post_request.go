package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUndoCheckoutUsingPostRequest Request Object
type ShowBatchUndoCheckoutUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o ShowBatchUndoCheckoutUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUndoCheckoutUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchUndoCheckoutUsingPostRequest", string(data)}, " ")
}
