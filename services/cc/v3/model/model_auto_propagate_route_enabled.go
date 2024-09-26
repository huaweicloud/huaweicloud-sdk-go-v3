package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoPropagateRouteEnabled struct {
}

func (o AutoPropagateRouteEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoPropagateRouteEnabled struct{}"
	}

	return strings.Join([]string{"AutoPropagateRouteEnabled", string(data)}, " ")
}
