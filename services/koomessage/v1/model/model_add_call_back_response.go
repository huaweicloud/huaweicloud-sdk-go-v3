package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCallBackResponse Response Object
type AddCallBackResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *Callback `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddCallBackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCallBackResponse struct{}"
	}

	return strings.Join([]string{"AddCallBackResponse", string(data)}, " ")
}
