package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkCreateUser **参数解释**：创建者信息，包含用户ID和用户名称。
type FlinkCreateUser struct {

	// **参数解释**：创建者用户ID，用于标识作业的创建者。 **取值范围**：长度为1~128个字符，只能由英文字母和数字组成。
	UserId string `json:"user_id"`

	// **参数解释**：创建者用户名称，用于显示作业的创建者名称。 **取值范围**：长度为0~128个字符。
	UserName string `json:"user_name"`
}

func (o FlinkCreateUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkCreateUser struct{}"
	}

	return strings.Join([]string{"FlinkCreateUser", string(data)}, " ")
}
