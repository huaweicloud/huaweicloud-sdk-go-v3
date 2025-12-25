package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAopworkflowRequest Request Object
type ImportAopworkflowRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *ImportAopworkflowRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportAopworkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAopworkflowRequest struct{}"
	}

	return strings.Join([]string{"ImportAopworkflowRequest", string(data)}, " ")
}
