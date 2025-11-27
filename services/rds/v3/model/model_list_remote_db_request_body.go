package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemoteDbRequestBody 查询远程SQL Server数据库请求参数。
type ListRemoteDbRequestBody struct {

	// 服务器ip。
	ServerIp string `json:"server_ip"`

	// 端口号。
	ServerPort string `json:"server_port"`

	// 登录名。
	LoginUserName string `json:"login_user_name"`

	// 登录密码。要求密码长度在5~64位之间。
	LoginUserPassword string `json:"login_user_password"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRemoteDbRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemoteDbRequestBody struct{}"
	}

	return strings.Join([]string{"ListRemoteDbRequestBody", string(data)}, " ")
}
