package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NonRequiredAutoPropagateRouteEnabled struct {
}

func (o NonRequiredAutoPropagateRouteEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonRequiredAutoPropagateRouteEnabled struct{}"
	}

	return strings.Join([]string{"NonRequiredAutoPropagateRouteEnabled", string(data)}, " ")
}
