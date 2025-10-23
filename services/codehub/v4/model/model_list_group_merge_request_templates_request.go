package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMergeRequestTemplatesRequest Request Object
type ListGroupMergeRequestTemplatesRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 搜索的模板名称 **取值范围：** 字符串长度不少于1，不超过100000。
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGroupMergeRequestTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMergeRequestTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListGroupMergeRequestTemplatesRequest", string(data)}, " ")
}
