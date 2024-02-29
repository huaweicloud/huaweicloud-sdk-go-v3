package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUndoCheckoutByAdminUsingPostRequest Request Object
type ShowUndoCheckoutByAdminUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o ShowUndoCheckoutByAdminUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUndoCheckoutByAdminUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowUndoCheckoutByAdminUsingPostRequest", string(data)}, " ")
}
