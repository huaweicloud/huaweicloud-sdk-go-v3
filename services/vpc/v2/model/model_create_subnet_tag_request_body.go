package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetTagRequestBody This is a auto create Body Object
type CreateSubnetTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateSubnetTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagRequestBody", string(data)}, " ")
}
