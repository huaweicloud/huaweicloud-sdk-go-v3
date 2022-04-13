package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KeystoneUserResultExtra struct {
	// IAM用户描述信息。

	Description *string `json:"description,omitempty"`
	// IAM用户密码状态。true：需要修改密码，false：正常。

	PwdStatus *bool `json:"pwd_status,omitempty"`
	// IAM用户退出系统前，在控制台最后访问的项目ID。

	LastProjectId *string `json:"last_project_id,omitempty"`
}

func (o KeystoneUserResultExtra) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUserResultExtra struct{}"
	}

	return strings.Join([]string{"KeystoneUserResultExtra", string(data)}, " ")
}
