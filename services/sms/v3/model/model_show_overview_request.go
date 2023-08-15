package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOverviewRequest Request Object
type ShowOverviewRequest struct {
}

func (o ShowOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowOverviewRequest", string(data)}, " ")
}
