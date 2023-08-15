package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCancelAuthorizationInterfaceResponse Response Object
type BatchCancelAuthorizationInterfaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCancelAuthorizationInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCancelAuthorizationInterfaceResponse struct{}"
	}

	return strings.Join([]string{"BatchCancelAuthorizationInterfaceResponse", string(data)}, " ")
}
