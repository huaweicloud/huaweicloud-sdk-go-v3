package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTableHistogramsRequestBody struct {

	// 毫秒时间戳
	From *int64 `json:"from,omitempty"`

	// 毫秒时间戳
	To *int64 `json:"to,omitempty"`

	// 检索查询条件, 语法介绍请参考帮助文档。
	Query string `json:"query"`
}

func (o ListTableHistogramsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableHistogramsRequestBody struct{}"
	}

	return strings.Join([]string{"ListTableHistogramsRequestBody", string(data)}, " ")
}
