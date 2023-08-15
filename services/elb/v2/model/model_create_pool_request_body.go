package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePoolRequestBody This is a auto create Body Object
type CreatePoolRequestBody struct {
	Pool *CreatePoolReq `json:"pool"`
}

func (o CreatePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePoolRequestBody", string(data)}, " ")
}
