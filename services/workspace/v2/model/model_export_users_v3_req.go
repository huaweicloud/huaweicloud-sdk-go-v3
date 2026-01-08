package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportUsersV3Req struct {

	// 桌面用户名，长度范围为1-20，不能包含特殊字符，不能以数字开头。支持模糊查询导出。
	UserName *string `json:"user_name,omitempty"`

	// 描述，支持模糊查询导出。
	Description *string `json:"description,omitempty"`

	// 激活类型，默认为用户激活。 * USER_ACTIVATE： 用户激活 * ADMIN_ACTIVATE： 管理员激活
	ActiveType *string `json:"active_type,omitempty"`

	// 语言，默认英文。 * zh_CN： 中文 * en_US： 英文
	Language *string `json:"language,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ExportUsersV3Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUsersV3Req struct{}"
	}

	return strings.Join([]string{"ExportUsersV3Req", string(data)}, " ")
}
