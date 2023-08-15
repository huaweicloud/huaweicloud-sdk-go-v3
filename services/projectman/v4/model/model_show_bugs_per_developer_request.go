package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBugsPerDeveloperRequest Request Object
type ShowBugsPerDeveloperRequest struct {

	// devcloud项目的32位id
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
