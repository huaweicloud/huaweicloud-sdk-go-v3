package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TriggerQpsDict HTTP限制列表
type TriggerQpsDict struct {

	// HTTP请求数分段ID，固定值为1
	HttpRequestPosId int64 `json:"http_request_pos_id"`

	// 每秒HTTP请求数（个/s）阈值
	HttpPacketPerSecond int64 `json:"http_packet_per_second"`
}

func (o TriggerQpsDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerQpsDict struct{}"
	}

	return strings.Join([]string{"TriggerQpsDict", string(data)}, " ")
}
