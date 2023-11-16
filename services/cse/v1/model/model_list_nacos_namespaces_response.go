package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNacosNamespacesResponse Response Object
type ListNacosNamespacesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ListNacosNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNacosNamespacesResponse struct{}"
	}

	return strings.Join([]string{"ListNacosNamespacesResponse", string(data)}, " ")
}
