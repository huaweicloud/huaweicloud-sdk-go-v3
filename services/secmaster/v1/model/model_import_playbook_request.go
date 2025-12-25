package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportPlaybookRequest Request Object
type ImportPlaybookRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *ImportPlaybookRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportPlaybookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportPlaybookRequest struct{}"
	}

	return strings.Join([]string{"ImportPlaybookRequest", string(data)}, " ")
}
