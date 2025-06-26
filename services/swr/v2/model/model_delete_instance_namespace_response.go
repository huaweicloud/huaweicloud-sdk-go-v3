package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceNamespaceResponse Response Object
type DeleteInstanceNamespaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceNamespaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceNamespaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceNamespaceResponse", string(data)}, " ")
}
