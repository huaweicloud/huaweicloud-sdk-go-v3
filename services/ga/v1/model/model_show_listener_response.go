package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowListenerResponse struct {
	Listener *ListenerDetail `json:"listener,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerResponse struct{}"
	}

	return strings.Join([]string{"ShowListenerResponse", string(data)}, " ")
}
