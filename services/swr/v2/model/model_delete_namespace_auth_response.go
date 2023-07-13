package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNamespaceAuthResponse Response Object
type DeleteNamespaceAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNamespaceAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNamespaceAuthResponse struct{}"
	}

	return strings.Join([]string{"DeleteNamespaceAuthResponse", string(data)}, " ")
}
