package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTagRequestBody", string(data)}, " ")
}
