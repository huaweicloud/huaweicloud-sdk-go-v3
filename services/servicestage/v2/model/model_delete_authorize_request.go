package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAuthorizeRequest struct {
	// 授权名称。

	Name string `json:"name"`
}

func (o DeleteAuthorizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuthorizeRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuthorizeRequest", string(data)}, " ")
}
