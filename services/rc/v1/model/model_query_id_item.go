package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryIdItem 查询ID项
type QueryIdItem struct {
}

func (o QueryIdItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryIdItem struct{}"
	}

	return strings.Join([]string{"QueryIdItem", string(data)}, " ")
}
