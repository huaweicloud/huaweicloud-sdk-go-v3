package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSBlackHoleEventRequest Request Object
type ListDDoSBlackHoleEventRequest struct {

	// 开始时间（毫秒时间戳）
	StartTime string `json:"start_time"`

	// 结束时间（毫秒时间戳）
	EndTime string `json:"end_time"`

	// 高防IP
	Ip *string `json:"ip,omitempty"`
}

func (o ListDDoSBlackHoleEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSBlackHoleEventRequest struct{}"
	}

	return strings.Join([]string{"ListDDoSBlackHoleEventRequest", string(data)}, " ")
}
