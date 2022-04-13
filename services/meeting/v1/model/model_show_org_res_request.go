package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOrgResRequest struct {
}

func (o ShowOrgResRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrgResRequest struct{}"
	}

	return strings.Join([]string{"ShowOrgResRequest", string(data)}, " ")
}
