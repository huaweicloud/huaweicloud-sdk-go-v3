package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAccessControlRequestBody struct {

	// 访问控制方式。 取值： - white：表示白名单。 - black：表示黑名单。
	Type string `json:"type"`

	// 控制访问的IP地址数组，最多可添加300个IP地址或网段。
	IpList []AccessControlRule `json:"ip_list"`
}

func (o CreateAccessControlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessControlRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAccessControlRequestBody", string(data)}, " ")
}
