package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CreateCredentialOption struct {
	// 待创建访问秘钥（AK/SK）的IAM用户ID，获取方式请参见：[获取账号、IAM用户、项目、用户组、委托的名称和ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	UserId string `json:"user_id"`
	// 访问密钥描述信息。

	Description *string `json:"description,omitempty"`
}

func (o CreateCredentialOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCredentialOption struct{}"
	}

	return strings.Join([]string{"CreateCredentialOption", string(data)}, " ")
}
