package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNewConfigsResponse struct {
	// 流量限制列表

	TrafficLimitedList *[]TriggerBpsDict `json:"traffic_limited_list,omitempty"`
	// HTTP限制列表

	HttpLimitedList *[]TriggerQpsDict `json:"http_limited_list,omitempty"`
	// 连接数限制列表

	ConnectionLimitedList *[]CleanLimitDict `json:"connection_limited_list,omitempty"`
	// 扩展配置列表

	ExtendDdosConfig *[]ExtendDDoSSet `json:"extend_ddos_config,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o ListNewConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNewConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListNewConfigsResponse", string(data)}, " ")
}
