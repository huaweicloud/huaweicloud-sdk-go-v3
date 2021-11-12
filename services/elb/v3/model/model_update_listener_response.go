package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateListenerResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Listener       *Listener `json:"listener,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerResponse struct{}"
	}

	return strings.Join([]string{"UpdateListenerResponse", string(data)}, " ")
}
