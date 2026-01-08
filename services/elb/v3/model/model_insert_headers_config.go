package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InsertHeadersConfig **参数解释**：要添加的请求头参数列表。  **约束限制**：不涉及
type InsertHeadersConfig struct {

	// **参数解释**：要添加请求头、响应头参数列表。  **约束限制**：不涉及
	Configs []InsertHeaderConfig `json:"configs"`
}

func (o InsertHeadersConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsertHeadersConfig struct{}"
	}

	return strings.Join([]string{"InsertHeadersConfig", string(data)}, " ")
}
