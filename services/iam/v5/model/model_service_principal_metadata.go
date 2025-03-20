package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServicePrincipalMetadata 服务主体信息。
type ServicePrincipalMetadata struct {

	// 服务主体，由\"service.\"开头，后跟一个长度为1到56个字符，只包含字母、数字和\"-\"的字符串。
	ServicePrincipal string `json:"service_principal"`

	// 云服务名称。
	ServiceCatalog string `json:"service_catalog"`

	// 用于显示的服务主体名称。
	DisplayName string `json:"display_name"`

	// 服务主体的描述。
	Description string `json:"description"`
}

func (o ServicePrincipalMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServicePrincipalMetadata struct{}"
	}

	return strings.Join([]string{"ServicePrincipalMetadata", string(data)}, " ")
}
