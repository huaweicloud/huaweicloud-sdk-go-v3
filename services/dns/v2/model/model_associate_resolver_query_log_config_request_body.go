package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateResolverQueryLogConfigRequestBody struct {
	Vpc *Vpc `json:"vpc"`
}

func (o AssociateResolverQueryLogConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResolverQueryLogConfigRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateResolverQueryLogConfigRequestBody", string(data)}, " ")
}
