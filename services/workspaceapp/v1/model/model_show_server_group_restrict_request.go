package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerGroupRestrictRequest Request Object
type ShowServerGroupRestrictRequest struct {
}

func (o ShowServerGroupRestrictRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupRestrictRequest struct{}"
	}

	return strings.Join([]string{"ShowServerGroupRestrictRequest", string(data)}, " ")
}
