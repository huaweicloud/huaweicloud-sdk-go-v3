package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDatabaseUserRequestBody struct {
	// 数据库用户名称。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、中划线、下划线和点。

	UserName string `json:"user_name"`
	// 用户所在的数据库。 - 长度为1~64位，可以包含大写字母（A~Z）、小写字母（a~z）、数字（0~9）、下划线。

	DbName string `json:"db_name"`
}

func (o DeleteDatabaseUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserRequestBody", string(data)}, " ")
}
