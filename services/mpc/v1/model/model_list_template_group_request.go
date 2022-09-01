package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplateGroupRequest struct {

	// 转码模板组ID，最多10个
	GroupId *[]string `json:"group_id,omitempty" xml:"group_id"`

	// 转码模板组名，最多10个
	GroupName *[]string `json:"group_name,omitempty" xml:"group_name"`

	// 分页编号。查询指定“group_id”时，该参数无效。  默认值：0。
	Page *int32 `json:"page,omitempty" xml:"page"`

	// 每页记录数。取值范围：[1,100]，指定group_id时该参数无效。
	Size *int32 `json:"size,omitempty" xml:"size"`
}

func (o ListTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupRequest", string(data)}, " ")
}
