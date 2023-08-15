package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVmsCallBackResponse Response Object
type AddVmsCallBackResponse struct {

	// 请求状态，固定200。
	Status *string `json:"status,omitempty"`

	// 状态描述。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddVmsCallBackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVmsCallBackResponse struct{}"
	}

	return strings.Join([]string{"AddVmsCallBackResponse", string(data)}, " ")
}
