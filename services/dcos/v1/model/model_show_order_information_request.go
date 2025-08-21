package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderInformationRequest Request Object
type ShowOrderInformationRequest struct {

	// 工单类型编码
	ModelCode string `json:"model_code"`
}

func (o ShowOrderInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderInformationRequest struct{}"
	}

	return strings.Join([]string{"ShowOrderInformationRequest", string(data)}, " ")
}
