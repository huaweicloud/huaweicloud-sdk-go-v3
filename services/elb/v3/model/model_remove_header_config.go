package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveHeaderConfig struct {

	// **参数解释**：被移除的请求头的参数名。  **取值范围**：1-40个字符，字母a-z（不区分大小写）、数字，短划线-和下划线_。
	Key string `json:"key"`
}

func (o RemoveHeaderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveHeaderConfig struct{}"
	}

	return strings.Join([]string{"RemoveHeaderConfig", string(data)}, " ")
}
