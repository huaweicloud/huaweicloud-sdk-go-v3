package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBugsPerDeveloperRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	Body *MetricRequest2 `json:"body,omitempty" xml:"body"`
}

func (o ShowBugsPerDeveloperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBugsPerDeveloperRequest struct{}"
	}

	return strings.Join([]string{"ShowBugsPerDeveloperRequest", string(data)}, " ")
}
