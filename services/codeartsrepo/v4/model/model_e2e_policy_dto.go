package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// E2ePolicyDto **参数解释：** E2E公共设置信息。
type E2ePolicyDto struct {

	// **参数解释：** 是否开启单号自动提取。 **取值范围：** true：开启单号自动提取，false：未开启单号自动提取。
	AutoExtract *bool `json:"auto_extract,omitempty"`

	// **参数解释：** 自动提取时的单号前缀合集，英文逗号分隔。 **取值范围：** 单个前缀长度最多200个字符，最多允许配置10个前缀。
	PrefixSymbol *string `json:"prefix_symbol,omitempty"`

	// **参数解释：** 自动提取时的分隔符。
	Separator *string `json:"separator,omitempty"`

	// **参数解释：** 自动提取时的单号后缀。
	SuffixSymbol *string `json:"suffix_symbol,omitempty"`
}

func (o E2ePolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "E2ePolicyDto struct{}"
	}

	return strings.Join([]string{"E2ePolicyDto", string(data)}, " ")
}
