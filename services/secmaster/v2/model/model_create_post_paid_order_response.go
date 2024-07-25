package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidOrderResponse Response Object
type CreatePostPaidOrderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePostPaidOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidOrderResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidOrderResponse", string(data)}, " ")
}
