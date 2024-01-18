package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartServerReq 批量删除服务器请求。
type BatchStartServerReq struct {

	// 批量请求的服务器ID列表，一次请求数量区间 [1, 20]。
	Items []string `json:"items"`
}

func (o BatchStartServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServerReq struct{}"
	}

	return strings.Join([]string{"BatchStartServerReq", string(data)}, " ")
}
