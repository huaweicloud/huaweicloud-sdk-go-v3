package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownUpTimeForSimCardReq struct {

	// 启用停用开关
	DownUpSwitch *int32 `json:"down_up_switch,omitempty"`

	// iccid，传入的sim_card_id为0,则根据iccid进行处理
	Iccid *string `json:"iccid,omitempty"`
}

func (o DownUpTimeForSimCardReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownUpTimeForSimCardReq struct{}"
	}

	return strings.Join([]string{"DownUpTimeForSimCardReq", string(data)}, " ")
}
