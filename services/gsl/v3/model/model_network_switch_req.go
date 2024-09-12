package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkSwitchReq struct {

	// 容器ID
	Cid *string `json:"cid,omitempty"`

	// 切换的目标网络
	CarrierType int32 `json:"carrier_type"`
}

func (o NetworkSwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkSwitchReq struct{}"
	}

	return strings.Join([]string{"NetworkSwitchReq", string(data)}, " ")
}
