package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubNetworkInterfaceResponse Response Object
type DeleteSubNetworkInterfaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubNetworkInterfaceResponse", string(data)}, " ")
}
