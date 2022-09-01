package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateListenerResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Listener       *Listener `json:"listener,omitempty" xml:"listener"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerResponse struct{}"
	}

	return strings.Join([]string{"CreateListenerResponse", string(data)}, " ")
}
