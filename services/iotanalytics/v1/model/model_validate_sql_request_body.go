package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidateSqlRequestBody struct {
	// 待执行的SQL语句。

	Sql *string `json:"sql,omitempty"`
}

func (o ValidateSqlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateSqlRequestBody struct{}"
	}

	return strings.Join([]string{"ValidateSqlRequestBody", string(data)}, " ")
}
