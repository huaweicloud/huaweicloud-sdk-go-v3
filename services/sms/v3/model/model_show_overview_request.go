package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOverviewRequest struct {
}

func (o ShowOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowOverviewRequest", string(data)}, " ")
}
