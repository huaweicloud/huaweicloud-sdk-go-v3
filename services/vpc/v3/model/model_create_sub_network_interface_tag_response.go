package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubNetworkInterfaceTagResponse Response Object
type CreateSubNetworkInterfaceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSubNetworkInterfaceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubNetworkInterfaceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateSubNetworkInterfaceTagResponse", string(data)}, " ")
}
