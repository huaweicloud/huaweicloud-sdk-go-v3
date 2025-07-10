package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeOverviewRequest Request Object
type ShowBestPracticeOverviewRequest struct {
}

func (o ShowBestPracticeOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeOverviewRequest", string(data)}, " ")
}
