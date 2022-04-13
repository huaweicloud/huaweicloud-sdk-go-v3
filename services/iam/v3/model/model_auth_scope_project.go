package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AuthScopeProject struct {
	// IAM用户所属账号的项目ID，获取方式请参见：[获取账号、IAM用户、项目、用户组、委托的名称和ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	Id *string `json:"id,omitempty"`
	// IAM用户所属账号的项目名称，获取方式请参见：[获取账号、IAM用户、项目、用户组、委托的名称和ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	Name *string `json:"name,omitempty"`
}

func (o AuthScopeProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthScopeProject struct{}"
	}

	return strings.Join([]string{"AuthScopeProject", string(data)}, " ")
}
