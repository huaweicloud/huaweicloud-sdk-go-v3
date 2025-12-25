package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchQuery 检索查询条件, 语法介绍请参考帮助文档。
type SearchQuery struct {
}

func (o SearchQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQuery struct{}"
	}

	return strings.Join([]string{"SearchQuery", string(data)}, " ")
}
