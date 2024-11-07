package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAsyncCommandsResponse Response Object
type ListAsyncCommandsResponse struct {

	// 设备队列命令列表。
	Commands *[]AsyncDeviceListCommand `json:"commands,omitempty"`

	Page           *QueueCommandPage `json:"page,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAsyncCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncCommandsResponse", string(data)}, " ")
}
