package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVersionAliasRequestBody struct {

	// 别名对应的版本名称。
	Version string `json:"version"`

	// 别名描述信息。
	Description *string `json:"description,omitempty"`

	// 灰度版本信息
	AdditionalVersionWeights map[string]int32 `json:"additional_version_weights,omitempty"`

	// 指定规则灰度策略信息
	AdditionalVersionStrategy map[string]VersionStrategy `json:"additional_version_strategy,omitempty"`
}

func (o UpdateVersionAliasRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVersionAliasRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVersionAliasRequestBody", string(data)}, " ")
}
