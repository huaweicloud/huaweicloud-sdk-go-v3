package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpcResponse Response Object
type ShowVpcResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	Vpc            *Vpc `json:"vpc,omitempty"`
	HttpStatusCode int  `json:"-"`
}

func (o ShowVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcResponse", string(data)}, " ")
}
