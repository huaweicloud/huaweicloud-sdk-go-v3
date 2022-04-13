package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateNamespaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateNamespaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNamespaceResponse struct{}"
	}

	return strings.Join([]string{"CreateNamespaceResponse", string(data)}, " ")
}
