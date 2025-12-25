package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCollectorParserRequest Request Object
type ImportCollectorParserRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ImportCollectorParserRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportCollectorParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCollectorParserRequest struct{}"
	}

	return strings.Join([]string{"ImportCollectorParserRequest", string(data)}, " ")
}
