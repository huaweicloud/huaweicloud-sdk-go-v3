package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowVpcResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Vpc            *Vpc `json:"vpc,omitempty" xml:"vpc"`
	HttpStatusCode int  `json:"-"`
}

func (o ShowVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcResponse", string(data)}, " ")
}
