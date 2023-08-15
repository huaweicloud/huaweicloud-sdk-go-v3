package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o ResourceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagRequestBody struct{}"
	}

	return strings.Join([]string{"ResourceTagRequestBody", string(data)}, " ")
}
