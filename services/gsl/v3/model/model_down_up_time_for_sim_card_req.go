package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownUpTimeForSimCardReq struct {

	// 启用停用开关
	DownUpSwitch *int32 `json:"down_up_switch,omitempty"`
}

func (o DownUpTimeForSimCardReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownUpTimeForSimCardReq struct{}"
	}

	return strings.Join([]string{"DownUpTimeForSimCardReq", string(data)}, " ")
}
