package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchSwitchoverResponse struct {
	// 批量主备倒换任务返回列表

	Results *[]SwitchoverResp `json:"results,omitempty"`
	// 总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchSwitchoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchoverResponse struct{}"
	}

	return strings.Join([]string{"BatchSwitchoverResponse", string(data)}, " ")
}
