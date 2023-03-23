package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PunishmentInfo struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 所属策略id
	Policyid *string `json:"policyid,omitempty"`

	// 拦截时间
	BlockTime *int32 `json:"block_time,omitempty"`

	// 攻击惩罚类别
	Category *string `json:"category,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o PunishmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PunishmentInfo struct{}"
	}

	return strings.Join([]string{"PunishmentInfo", string(data)}, " ")
}
