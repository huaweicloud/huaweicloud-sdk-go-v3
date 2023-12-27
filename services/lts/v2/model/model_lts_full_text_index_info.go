package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsFullTextIndexInfo struct {

	// 是否开启全文索引
	Enable bool `json:"enable"`

	// 是否大小写敏感
	CaseSensitive bool `json:"caseSensitive"`

	// 是否包含中文
	IncludeChinese bool `json:"includeChinese"`

	// 自定义分词符
	Tokenizer string `json:"tokenizer"`

	// 特殊分词符
	Ascii *[]string `json:"ascii,omitempty"`
}

func (o LtsFullTextIndexInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsFullTextIndexInfo struct{}"
	}

	return strings.Join([]string{"LtsFullTextIndexInfo", string(data)}, " ")
}
