package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectivityTestRequestBody 备份导入页面实例连接测试请求
type CreateConnectivityTestRequestBody struct {

	// 登录密码
	Password *string `json:"password,omitempty"`
}

func (o CreateConnectivityTestRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectivityTestRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConnectivityTestRequestBody", string(data)}, " ")
}
