package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupMergeRequestTemplateRequest Request Object
type CreateGroupMergeRequestTemplateRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *CreateMergeRequestTemplateDto `json:"body,omitempty"`
}

func (o CreateGroupMergeRequestTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupMergeRequestTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateGroupMergeRequestTemplateRequest", string(data)}, " ")
}
