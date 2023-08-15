package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNamespacesResponse Response Object
type DeleteNamespacesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNamespacesResponse struct{}"
	}

	return strings.Join([]string{"DeleteNamespacesResponse", string(data)}, " ")
}
