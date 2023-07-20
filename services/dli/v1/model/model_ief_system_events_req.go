package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IefSystemEventsReq IEF系统事件的请求body体。
type IefSystemEventsReq struct {
	Data *IefEvents `json:"data"`
}

func (o IefSystemEventsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IefSystemEventsReq struct{}"
	}

	return strings.Join([]string{"IefSystemEventsReq", string(data)}, " ")
}
