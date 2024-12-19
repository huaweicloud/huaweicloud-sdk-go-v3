package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RouteNextHop 下一跳id
type RouteNextHop struct {
}

func (o RouteNextHop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteNextHop struct{}"
	}

	return strings.Join([]string{"RouteNextHop", string(data)}, " ")
}
