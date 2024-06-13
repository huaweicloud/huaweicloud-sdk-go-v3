package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityGroupTagRequestBody This is a auto create Body Object
type CreateSecurityGroupTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateSecurityGroupTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupTagRequestBody", string(data)}, " ")
}
