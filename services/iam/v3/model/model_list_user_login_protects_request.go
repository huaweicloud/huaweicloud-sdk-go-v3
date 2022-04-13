package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListUserLoginProtectsRequest struct {
}

func (o ListUserLoginProtectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserLoginProtectsRequest struct{}"
	}

	return strings.Join([]string{"ListUserLoginProtectsRequest", string(data)}, " ")
}
