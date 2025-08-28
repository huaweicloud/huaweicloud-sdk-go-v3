package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecycleBinRequestBody **参数解释**：更新回收站启用开关的请求体。
type RecycleBinRequestBody struct {

	// **参数解释**：是否启用回收站。  **取值范围**： - true：启用回收站。 - false：不启用回收站。
	Enable *bool `json:"enable,omitempty"`
}

func (o RecycleBinRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinRequestBody struct{}"
	}

	return strings.Join([]string{"RecycleBinRequestBody", string(data)}, " ")
}
