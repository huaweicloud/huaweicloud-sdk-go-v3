package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeShareChargeModeRequestBody 按需转包周期请求体参数
type ChangeShareChargeModeRequestBody struct {
	BssParam *BssInfo `json:"bss_param"`
}

func (o ChangeShareChargeModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareChargeModeRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeShareChargeModeRequestBody", string(data)}, " ")
}
