package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourceTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateResourceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceTagRequestBody", string(data)}, " ")
}
