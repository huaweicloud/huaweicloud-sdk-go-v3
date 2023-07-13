package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNamespaceAuthResponse Response Object
type CreateNamespaceAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateNamespaceAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNamespaceAuthResponse struct{}"
	}

	return strings.Join([]string{"CreateNamespaceAuthResponse", string(data)}, " ")
}
