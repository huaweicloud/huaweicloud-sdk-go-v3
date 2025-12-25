package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorParserRequest Request Object
type CreateCollectorParserRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateParserDto `json:"body,omitempty"`
}

func (o CreateCollectorParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorParserRequest struct{}"
	}

	return strings.Join([]string{"CreateCollectorParserRequest", string(data)}, " ")
}
