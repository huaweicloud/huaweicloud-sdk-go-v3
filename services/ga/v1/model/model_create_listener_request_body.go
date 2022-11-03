package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// create Listener request
type CreateListenerRequestBody struct {
	Listener *CreateListenerOption `json:"listener"`
}

func (o CreateListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateListenerRequestBody", string(data)}, " ")
}
