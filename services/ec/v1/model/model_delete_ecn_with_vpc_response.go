package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEcnWithVpcResponse Response Object
type DeleteEcnWithVpcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEcnWithVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEcnWithVpcResponse struct{}"
	}

	return strings.Join([]string{"DeleteEcnWithVpcResponse", string(data)}, " ")
}
