package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNacosNamespacesResponse Response Object
type UpdateNacosNamespacesResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNacosNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNacosNamespacesResponse struct{}"
	}

	return strings.Join([]string{"UpdateNacosNamespacesResponse", string(data)}, " ")
}
