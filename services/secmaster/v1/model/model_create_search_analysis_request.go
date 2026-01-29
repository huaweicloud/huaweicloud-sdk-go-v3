package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchAnalysisRequest Request Object
type CreateSearchAnalysisRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateAnalysisRequestBody `json:"body,omitempty"`
}

func (o CreateSearchAnalysisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchAnalysisRequest struct{}"
	}

	return strings.Join([]string{"CreateSearchAnalysisRequest", string(data)}, " ")
}
