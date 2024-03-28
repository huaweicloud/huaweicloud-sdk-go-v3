package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalVariable 全局变量
type GlobalVariable struct {

	// 全局变量ID
	Id *int64 `json:"id,omitempty"`

	// 变量名称
	VarName *string `json:"var_name,omitempty"`

	// 变量的值
	VarValue *string `json:"var_value,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 是否为敏感变量
	IsSensitive *bool `json:"is_sensitive,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 创建时间。为UTC的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间。为UTC的时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o GlobalVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalVariable struct{}"
	}

	return strings.Join([]string{"GlobalVariable", string(data)}, " ")
}
