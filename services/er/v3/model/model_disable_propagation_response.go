package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisablePropagationResponse Response Object
type DisablePropagationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisablePropagationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisablePropagationResponse struct{}"
	}

	return strings.Join([]string{"DisablePropagationResponse", string(data)}, " ")
}
