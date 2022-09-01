package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteProcessReqBody struct {

	// 数据库用户ID
	DbUserId string `json:"db_user_id" xml:"db_user_id"`

	// 会话ID列表。process_ids、user、database至少指定一个参数。
	ProcessIds *[]string `json:"process_ids,omitempty" xml:"process_ids"`

	// 用户
	User *string `json:"user,omitempty" xml:"user"`

	// 数据库名称
	Database *string `json:"database,omitempty" xml:"database"`
}

func (o DeleteProcessReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProcessReqBody struct{}"
	}

	return strings.Join([]string{"DeleteProcessReqBody", string(data)}, " ")
}
