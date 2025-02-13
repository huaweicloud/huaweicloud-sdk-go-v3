package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNacosNamespacesResponse Response Object
type CreateNacosNamespacesResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNacosNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNacosNamespacesResponse struct{}"
	}

	return strings.Join([]string{"CreateNacosNamespacesResponse", string(data)}, " ")
}
