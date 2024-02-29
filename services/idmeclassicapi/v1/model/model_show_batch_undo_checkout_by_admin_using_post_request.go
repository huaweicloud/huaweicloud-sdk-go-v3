package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUndoCheckoutByAdminUsingPostRequest Request Object
type ShowBatchUndoCheckoutByAdminUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o ShowBatchUndoCheckoutByAdminUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUndoCheckoutByAdminUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchUndoCheckoutByAdminUsingPostRequest", string(data)}, " ")
}
