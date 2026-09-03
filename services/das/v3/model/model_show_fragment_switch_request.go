package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFragmentSwitchRequest Request Object
type ShowFragmentSwitchRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`
}

func (o ShowFragmentSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFragmentSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowFragmentSwitchRequest", string(data)}, " ")
}
