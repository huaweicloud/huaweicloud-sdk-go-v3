package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportLogoResponse Response Object
type ShowSecurityReportLogoResponse struct {

	// **参数解释**： logo文件base64内容 **取值范围**： 字符长度32-36位
	LogoContent *string `json:"logo_content,omitempty"`

	// **参数解释**： logo文件名称 **取值范围**： 字符长度0-128位
	LogoName       *string `json:"logo_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecurityReportLogoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportLogoResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportLogoResponse", string(data)}, " ")
}
