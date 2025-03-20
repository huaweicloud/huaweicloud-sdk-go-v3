package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServiceLinkedAgencyReqBody struct {

	// 服务主体，由\"service.\"开头，后跟一个长度为1到56个字符，只包含字母、数字和\"-\"的字符串。
	ServicePrincipal string `json:"service_principal"`

	// 服务关联委托描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`
}

func (o CreateServiceLinkedAgencyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceLinkedAgencyReqBody struct{}"
	}

	return strings.Join([]string{"CreateServiceLinkedAgencyReqBody", string(data)}, " ")
}
