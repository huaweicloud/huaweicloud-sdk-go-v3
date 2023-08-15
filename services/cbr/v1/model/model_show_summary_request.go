package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSummaryRequest Request Object
type ShowSummaryRequest struct {
}

func (o ShowSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowSummaryRequest", string(data)}, " ")
}
