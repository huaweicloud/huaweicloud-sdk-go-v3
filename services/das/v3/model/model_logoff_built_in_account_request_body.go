package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogoffBuiltInAccountRequestBody struct {

	// 数据库引擎，仅支持MySQL
	DatastoreType string `json:"datastore_type"`
}

func (o LogoffBuiltInAccountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffBuiltInAccountRequestBody struct{}"
	}

	return strings.Join([]string{"LogoffBuiltInAccountRequestBody", string(data)}, " ")
}
