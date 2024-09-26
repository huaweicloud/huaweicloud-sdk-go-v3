package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NonDefaultAutoPropagateRouteEnabled struct {
}

func (o NonDefaultAutoPropagateRouteEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonDefaultAutoPropagateRouteEnabled struct{}"
	}

	return strings.Join([]string{"NonDefaultAutoPropagateRouteEnabled", string(data)}, " ")
}
