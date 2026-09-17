package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnalysisResultRequest Request Object
type ListAnalysisResultRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *QueryAnalysisResultBody `json:"body,omitempty"`
}

func (o ListAnalysisResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnalysisResultRequest struct{}"
	}

	return strings.Join([]string{"ListAnalysisResultRequest", string(data)}, " ")
}
