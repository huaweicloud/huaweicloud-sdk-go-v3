package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Query 查询语句
type Query struct {
}

func (o Query) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Query struct{}"
	}

	return strings.Join([]string{"Query", string(data)}, " ")
}
