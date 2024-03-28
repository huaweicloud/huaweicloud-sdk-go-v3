package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefSystemEventsRequestBody IEF系统事件的请求body体。
type CreateIefSystemEventsRequestBody struct {
	Data *IefEvent `json:"data"`
}

func (o CreateIefSystemEventsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefSystemEventsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIefSystemEventsRequestBody", string(data)}, " ")
}
