package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSystemEventDetailRequest Request Object
type ShowSystemEventDetailRequest struct {

	// 系统订阅名称
	EventId string `json:"event_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ShowSystemEventDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSystemEventDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSystemEventDetailRequest", string(data)}, " ")
}
