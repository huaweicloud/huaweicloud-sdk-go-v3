package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePostPaidInstanceRequestBody struct {
	Instance *CreatePostpaidInstanceOption `json:"instance"`
}

func (o CreatePostPaidInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceRequestBody", string(data)}, " ")
}
