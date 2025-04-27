package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateOrDisassociateRouterWithRuleRequestBody struct {
	Router *RouterForRule `json:"router"`
}

func (o AssociateOrDisassociateRouterWithRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateOrDisassociateRouterWithRuleRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateOrDisassociateRouterWithRuleRequestBody", string(data)}, " ")
}
