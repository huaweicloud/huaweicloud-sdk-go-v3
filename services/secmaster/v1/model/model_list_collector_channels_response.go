package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorChannelsResponse Response Object
type ListCollectorChannelsResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 相关描述信息
	Records        *[]Channel `json:"records,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListCollectorChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorChannelsResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorChannelsResponse", string(data)}, " ")
}
