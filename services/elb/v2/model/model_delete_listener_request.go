package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteListenerRequest struct {

	// （不再支持）级联删除负载均衡器
	Cascade *bool `json:"cascade,omitempty"`

	// 监听器id
	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerRequest", string(data)}, " ")
}
