package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorInstancesResponse Response Object
type ListCollectorInstancesResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 相关描述信息
	Records        *[]ChannelInstance `json:"records,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCollectorInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorInstancesResponse", string(data)}, " ")
}
