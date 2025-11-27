package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbmSuccessResponseBody 成功响应体。
type DbmSuccessResponseBody struct {
}

func (o DbmSuccessResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbmSuccessResponseBody struct{}"
	}

	return strings.Join([]string{"DbmSuccessResponseBody", string(data)}, " ")
}
