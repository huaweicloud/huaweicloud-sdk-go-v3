package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVersionAliasRequestBody 创建别名请求体
type CreateVersionAliasRequestBody struct {

	// 别名名称。
	Name string `json:"name"`

	// 别名对应的版本名称。
	Version string `json:"version"`

	// 别名描述信息。
	Description *string `json:"description,omitempty"`

	// 百分比灰度配置信息
	AdditionalVersionWeights map[string]int32 `json:"additional_version_weights,omitempty"`

	// 指定规则灰度策略信息
	AdditionalVersionStrategy map[string]VersionStrategy `json:"additional_version_strategy,omitempty"`
}

func (o CreateVersionAliasRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVersionAliasRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVersionAliasRequestBody", string(data)}, " ")
}
