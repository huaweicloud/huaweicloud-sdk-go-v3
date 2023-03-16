package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDeleteDatabaseUserRequest struct {

	// 数据库用户名。
	Name string `json:"name"`

	// 主机地址
	Host string `json:"host"`
}

func (o ListDeleteDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeleteDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"ListDeleteDatabaseUserRequest", string(data)}, " ")
}
