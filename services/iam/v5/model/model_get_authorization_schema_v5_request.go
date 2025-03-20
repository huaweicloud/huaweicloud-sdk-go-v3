package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAuthorizationSchemaV5Request Request Object
type GetAuthorizationSchemaV5Request struct {

	// 服务名称缩写，长度为1到56个字符，只包含字母、数字和\"-\"的字符串。
	ServiceCode string `json:"service_code"`
}

func (o GetAuthorizationSchemaV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAuthorizationSchemaV5Request struct{}"
	}

	return strings.Join([]string{"GetAuthorizationSchemaV5Request", string(data)}, " ")
}
