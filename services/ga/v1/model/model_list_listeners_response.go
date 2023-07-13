package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListListenersResponse Response Object
type ListListenersResponse struct {

	// 监听器列表。
	Listeners *[]ListenerDetail `json:"listeners,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListListenersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersResponse struct{}"
	}

	return strings.Join([]string{"ListListenersResponse", string(data)}, " ")
}
