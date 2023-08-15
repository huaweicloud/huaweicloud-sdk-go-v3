package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcResponse Response Object
type DeleteVpcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcResponse", string(data)}, " ")
}
