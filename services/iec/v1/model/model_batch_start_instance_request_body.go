package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量启动边缘实例请求体。
type BatchStartInstanceRequestBody struct {
	OsStart *BatchStart `json:"os-start,omitempty"`
}

func (o BatchStartInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStartInstanceRequestBody", string(data)}, " ")
}
