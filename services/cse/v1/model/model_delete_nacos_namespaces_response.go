package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNacosNamespacesResponse Response Object
type DeleteNacosNamespacesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNacosNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNacosNamespacesResponse struct{}"
	}

	return strings.Join([]string{"DeleteNacosNamespacesResponse", string(data)}, " ")
}
