package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIssuesWrokFlowConfigResponse struct {

	// 流转数据
	Workflows      *[]ScrumStatusFlowVo `json:"workflows,omitempty" xml:"workflows"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowIssuesWrokFlowConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssuesWrokFlowConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowIssuesWrokFlowConfigResponse", string(data)}, " ")
}
