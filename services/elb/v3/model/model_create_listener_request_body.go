package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateListenerRequestBody This is a auto create Body Object
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
