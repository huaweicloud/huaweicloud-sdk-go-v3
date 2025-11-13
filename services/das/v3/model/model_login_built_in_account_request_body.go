package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginBuiltInAccountRequestBody 请求体
type LoginBuiltInAccountRequestBody struct {

	// 数据库引擎，仅支持MySQL
	DatastoreType string `json:"datastore_type"`
}

func (o LoginBuiltInAccountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginBuiltInAccountRequestBody struct{}"
	}

	return strings.Join([]string{"LoginBuiltInAccountRequestBody", string(data)}, " ")
}
