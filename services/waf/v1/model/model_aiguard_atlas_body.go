package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardAtlasBody struct {

	// 图库名称，同project唯一
	Name string `json:"name"`

	// 图库类型 0=白库（豁免/可信样本） / 1=黑库（敏感/拦截样本）
	Type int32 `json:"type"`

	// 图库描述
	Description *string `json:"description,omitempty"`
}

func (o AiguardAtlasBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardAtlasBody struct{}"
	}

	return strings.Join([]string{"AiguardAtlasBody", string(data)}, " ")
}
