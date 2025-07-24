package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteActiveDirectoryDomainRequestBody AD域配置信息
type DeleteActiveDirectoryDomainRequestBody struct {

	// 服务账号，在创建域服务器时指定，一般默认为administrator
	ServiceAccount string `json:"service_account"`

	// 账号对应密码
	Password string `json:"password"`
}

func (o DeleteActiveDirectoryDomainRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActiveDirectoryDomainRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteActiveDirectoryDomainRequestBody", string(data)}, " ")
}
