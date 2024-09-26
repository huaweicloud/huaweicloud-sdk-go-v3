package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NonRequiredAutoAssociateRouteEnabled struct {
}

func (o NonRequiredAutoAssociateRouteEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonRequiredAutoAssociateRouteEnabled struct{}"
	}

	return strings.Join([]string{"NonRequiredAutoAssociateRouteEnabled", string(data)}, " ")
}
