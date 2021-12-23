package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNamespaceAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNamespaceAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNamespaceAuthResponse struct{}"
	}

	return strings.Join([]string{"UpdateNamespaceAuthResponse", string(data)}, " ")
}
