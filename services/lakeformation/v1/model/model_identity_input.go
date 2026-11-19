package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdentityInput 用户名称，用户组，角色及OBS路径信息
type IdentityInput struct {

	// IAM用户
	UserNames *[]string `json:"user_names,omitempty"`

	// 用户组
	Groups *[]string `json:"groups,omitempty"`

	// 角色
	Roles *[]string `json:"roles,omitempty"`

	// OBS路径（可选），用于指定获取临时AKSK访问的对象路径，仅支持并行文件系统（PFS）桶，不支持对象桶。格式：obs://桶名/对象键
	Location *string `json:"location,omitempty"`
}

func (o IdentityInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityInput struct{}"
	}

	return strings.Join([]string{"IdentityInput", string(data)}, " ")
}
