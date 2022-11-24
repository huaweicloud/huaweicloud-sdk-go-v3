package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 需要关联的Router(VPC)。
type AssociateRouterReq struct {
	Router *Router `json:"router"`
}

func (o AssociateRouterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouterReq struct{}"
	}

	return strings.Join([]string{"AssociateRouterReq", string(data)}, " ")
}
