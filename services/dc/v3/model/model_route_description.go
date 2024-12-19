package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RouteDescription 路由描述
type RouteDescription struct {
}

func (o RouteDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteDescription struct{}"
	}

	return strings.Join([]string{"RouteDescription", string(data)}, " ")
}
