package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAsyncHistoryCommandsResponse Response Object
type ListAsyncHistoryCommandsResponse struct {

	// 设备历史命令列表。
	Commands *[]AsyncDeviceListCommand `json:"commands,omitempty"`

	Page           *HistoryCommandPage `json:"page,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAsyncHistoryCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncHistoryCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncHistoryCommandsResponse", string(data)}, " ")
}
