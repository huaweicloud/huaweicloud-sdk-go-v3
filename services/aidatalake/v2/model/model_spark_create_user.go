package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkCreateUser **参数解释**：创建作业的用户信息。
type SparkCreateUser struct {

	// **参数解释**：创建作业的用户ID，用于标识作业创建者。 **取值范围**：长度为1~256个字符。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：创建作业的用户名称，用于标识作业创建者。 **取值范围**：长度为1~256个字符。
	UserName *string `json:"user_name,omitempty"`
}

func (o SparkCreateUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkCreateUser struct{}"
	}

	return strings.Join([]string{"SparkCreateUser", string(data)}, " ")
}
