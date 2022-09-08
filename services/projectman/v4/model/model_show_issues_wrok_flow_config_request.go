package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowIssuesWrokFlowConfigRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项类型id [2,3,4,5,6,7]
	TrackerId int32 `json:"tracker_id"`
}

func (o ShowIssuesWrokFlowConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssuesWrokFlowConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowIssuesWrokFlowConfigRequest", string(data)}, " ")
}
