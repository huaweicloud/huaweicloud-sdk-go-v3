package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 版本别名结构
type CreateVersionAliasRequestBody struct {

	// 要获取的别名名称。
	Name string `json:"name" xml:"name"`

	// 别名对应的版本名称。
	Version string `json:"version" xml:"version"`

	// 别名描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 灰度版本信息
	AdditionalVersionWeights map[string]int32 `json:"additional_version_weights,omitempty" xml:"additional_version_weights"`
}

func (o CreateVersionAliasRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVersionAliasRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVersionAliasRequestBody", string(data)}, " ")
}
