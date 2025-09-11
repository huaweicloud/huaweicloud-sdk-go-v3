package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Filter struct {

	// 过滤类型，可选name，tag
	Type string `json:"type"`

	// 过滤类型对应的正则表达式(type为name时，对应的是命名空间及制品仓库，例如library/_**； type为tag时，对应的是制品版本，例如：**repo和**tag； 正则表达式有多个时，用英文逗号分隔，且在最外层加{}，如library/{test,test*,*test})
	Value string `json:"value"`
}

func (o Filter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Filter struct{}"
	}

	return strings.Join([]string{"Filter", string(data)}, " ")
}
