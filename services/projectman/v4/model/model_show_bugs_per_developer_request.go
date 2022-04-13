package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBugsPerDeveloperRequest struct {
	// devcloud的项目ID

	ProjectId string `json:"project_id"`

	Body *MetricRequest2 `json:"body,omitempty"`
}

func (o ShowBugsPerDeveloperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBugsPerDeveloperRequest struct{}"
	}

	return strings.Join([]string{"ShowBugsPerDeveloperRequest", string(data)}, " ")
}
