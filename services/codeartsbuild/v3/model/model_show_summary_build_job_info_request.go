package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSummaryBuildJobInfoRequest Request Object
type ShowSummaryBuildJobInfoRequest struct {
}

func (o ShowSummaryBuildJobInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSummaryBuildJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSummaryBuildJobInfoRequest", string(data)}, " ")
}
