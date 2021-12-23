package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateStreamResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamResponse struct{}"
	}

	return strings.Join([]string{"CreateStreamResponse", string(data)}, " ")
}
