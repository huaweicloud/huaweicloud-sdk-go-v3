package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBugDensityV2Request struct {

	// devcloud的项目ID
	ProjectId string `json:"project_id"`

	Body *MetricRequestV2 `json:"body,omitempty"`
}

func (o ShowBugDensityV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBugDensityV2Request struct{}"
	}

	return strings.Join([]string{"ShowBugDensityV2Request", string(data)}, " ")
}
