package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScrumIssueSeveritiesRequest Request Object
type ShowScrumIssueSeveritiesRequest struct {
}

func (o ShowScrumIssueSeveritiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScrumIssueSeveritiesRequest struct{}"
	}

	return strings.Join([]string{"ShowScrumIssueSeveritiesRequest", string(data)}, " ")
}
