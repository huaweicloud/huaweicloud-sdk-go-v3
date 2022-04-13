package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 共享带宽插入/移除弹性公网IP的publicip_info字段
type InsertPublicipInfo struct {
	// 功能说明：若publicip_id为弹性公网IP的id，则该字段可自动忽略。若publicip_id为IPv6端口PORT的id，则该字段必填：5_dualStack(目前仅北京4局点支持)

	PublicipType *string `json:"publicip_type,omitempty"`
	// 功能说明：带宽对应的弹性公网IP或IPv6端口PORT的唯一标识

	PublicipId string `json:"publicip_id"`
}

func (o InsertPublicipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsertPublicipInfo struct{}"
	}

	return strings.Join([]string{"InsertPublicipInfo", string(data)}, " ")
}
