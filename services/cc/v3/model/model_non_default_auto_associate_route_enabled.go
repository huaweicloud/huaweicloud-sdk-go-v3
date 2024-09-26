package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NonDefaultAutoAssociateRouteEnabled struct {
}

func (o NonDefaultAutoAssociateRouteEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonDefaultAutoAssociateRouteEnabled struct{}"
	}

	return strings.Join([]string{"NonDefaultAutoAssociateRouteEnabled", string(data)}, " ")
}
