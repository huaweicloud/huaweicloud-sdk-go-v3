package model

import (
	"encoding/json"

	"strings"
)

type DeleteProcessReqBody struct {
	// 数据库用户ID

	DbUserId string `json:"db_user_id"`
	// 会话ID列表。process_ids、user、database至少指定一个参数。

	ProcessIds *[]string `json:"process_ids,omitempty"`
	// 用户

	User *string `json:"user,omitempty"`
	// 数据库名称

	Database *string `json:"database,omitempty"`
}

func (o DeleteProcessReqBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProcessReqBody struct{}"
	}

	return strings.Join([]string{"DeleteProcessReqBody", string(data)}, " ")
}
