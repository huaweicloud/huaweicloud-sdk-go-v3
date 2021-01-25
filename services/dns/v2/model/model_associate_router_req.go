package model

import (
	"encoding/json"

	"strings"
)

// 需要关联的Router(VPC)。
type AssociateRouterReq struct {
	Router *Router `json:"router"`
}

func (o AssociateRouterReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateRouterReq struct{}"
	}

	return strings.Join([]string{"AssociateRouterReq", string(data)}, " ")
}
