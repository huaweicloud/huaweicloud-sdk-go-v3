package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserData 创建实例过程中待注入实例自定义数据。支持注入文本、文本文件。   说明： user_data的值为base64编码之后的内容。 注入内容（编码之前的内容）最大长度为32K。
type UserData struct {
}

func (o UserData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserData struct{}"
	}

	return strings.Join([]string{"UserData", string(data)}, " ")
}
