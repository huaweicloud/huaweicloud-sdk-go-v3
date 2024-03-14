package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateActiveCodeReq 修改激活码请求。
type UpdateActiveCodeReq struct {

	// 有效天数（0表示长期有效）。
	ValidPeriod *int32 `json:"valid_period,omitempty"`
}

func (o UpdateActiveCodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateActiveCodeReq struct{}"
	}

	return strings.Join([]string{"UpdateActiveCodeReq", string(data)}, " ")
}
