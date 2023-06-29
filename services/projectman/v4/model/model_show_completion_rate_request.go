package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCompletionRateRequest Request Object
type ShowCompletionRateRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *MetricRequest3 `json:"body,omitempty"`
}

func (o ShowCompletionRateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompletionRateRequest struct{}"
	}

	return strings.Join([]string{"ShowCompletionRateRequest", string(data)}, " ")
}
