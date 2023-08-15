package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcResponse Response Object
type CreateVpcResponse struct {
	Vpc            *Vpc `json:"vpc,omitempty"`
	HttpStatusCode int  `json:"-"`
}

func (o CreateVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcResponse", string(data)}, " ")
}
