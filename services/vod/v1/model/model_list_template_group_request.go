package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplateGroupRequest struct {

	// 模板组id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 模板启用状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 分页编号。  默认为0，指定group_id时该参数无效。
	Page *int32 `json:"page,omitempty" xml:"page"`

	// 每页记录数。  默认为10，范围[1,100]。指定group_id时该参数无效。
	Size *int32 `json:"size,omitempty" xml:"size"`
}

func (o ListTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupRequest", string(data)}, " ")
}
