package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListListenersResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`
	// listener的列表。

	Listeners      *[]Listener `json:"listeners,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListListenersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersResponse struct{}"
	}

	return strings.Join([]string{"ListListenersResponse", string(data)}, " ")
}
