package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RemoveVpcExtendCidrResponse struct {
	Vpc *Vpc `json:"vpc,omitempty" xml:"vpc"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoveVpcExtendCidrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVpcExtendCidrResponse struct{}"
	}

	return strings.Join([]string{"RemoveVpcExtendCidrResponse", string(data)}, " ")
}
