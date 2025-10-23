package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupMergeRequestTemplateRequest Request Object
type UpdateGroupMergeRequestTemplateRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 合并请求模板id。
	TemplateId int32 `json:"template_id"`

	Body *CreateMergeRequestTemplateDto `json:"body,omitempty"`
}

func (o UpdateGroupMergeRequestTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupMergeRequestTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupMergeRequestTemplateRequest", string(data)}, " ")
}
