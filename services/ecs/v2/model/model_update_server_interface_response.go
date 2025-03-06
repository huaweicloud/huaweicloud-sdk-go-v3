package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerInterfaceResponse Response Object
type UpdateServerInterfaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerInterfaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerInterfaceResponse", string(data)}, " ")
}
