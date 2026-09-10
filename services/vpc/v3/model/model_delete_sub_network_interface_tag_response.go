package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubNetworkInterfaceTagResponse Response Object
type DeleteSubNetworkInterfaceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubNetworkInterfaceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubNetworkInterfaceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubNetworkInterfaceTagResponse", string(data)}, " ")
}
