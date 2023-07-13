package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaurusSwitchoverRequest struct {

	// 只读节点ID，倒换后为主节点。
	NodeId string `json:"node_id"`
}

func (o TaurusSwitchoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusSwitchoverRequest struct{}"
	}

	return strings.Join([]string{"TaurusSwitchoverRequest", string(data)}, " ")
}
