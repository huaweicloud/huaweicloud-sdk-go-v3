package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtDataSourceReq 数据源请求
type ExtDataSourceReq struct {

	// 数据源id
	DataSourceId *string `json:"data_source_id,omitempty"`

	// 类型
	Type string `json:"type"`

	// 数据源名称
	DataSourceName string `json:"data_source_name"`

	// 用户名
	UserName string `json:"user_name"`

	// 密码
	UserPwd *string `json:"user_pwd,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 重启
	Reboot *bool `json:"reboot,omitempty"`

	// 数据库
	ConnectInfo *string `json:"connect_info,omitempty"`
}

func (o ExtDataSourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtDataSourceReq struct{}"
	}

	return strings.Join([]string{"ExtDataSourceReq", string(data)}, " ")
}
