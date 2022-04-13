package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 邀请企业内用户的DTO对象
type AddUserRequestBody struct {
	// 用户名称

	Name string `json:"name"`
	// 后台自动识别是手机还是邮箱,若为手机号，则要求和国家码匹配

	Contact string `json:"contact"`
	// 国家码。 国家和国家码的对应关系请参考:https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
	// 部门编号，若不携带则默认根部门

	DeptCode *string `json:"deptCode,omitempty"`
	// 职位

	Title *string `json:"title,omitempty"`
	// 通讯录排序等级，序号越低优先级越高

	SortLevel *int32 `json:"sortLevel,omitempty"`
	// 备注

	Desc *string `json:"desc,omitempty"`
}

func (o AddUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserRequestBody struct{}"
	}

	return strings.Join([]string{"AddUserRequestBody", string(data)}, " ")
}
