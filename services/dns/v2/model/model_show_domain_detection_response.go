package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainDetectionResponse Response Object
type ShowDomainDetectionResponse struct {

	// 待诊断记录集的名称。
	DomainName *string `json:"domain_name,omitempty"`

	// 待诊断记录集的类型。 取值范围：CNAME、TXT、MX。
	Type *string `json:"type,omitempty"`

	// 域名诊断状态。  取值范围：  OK：解析成功 FAILED：whois查询失败 NOT_REGISTERED：通过whois查询，域名未注册 CANNOT_RESOLVE：通过whois查询，域名无法解析 NOT_HWDNS：未托管在华为云 NO_WEBSITE_RECORD：未配置网站解析记录 NO_EMAIL_RECORD：未配置邮箱解析记录 NO_DEFAULT_VIEW：未配置默认解析 NOT_EFFECT：权威服务器未生效
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDomainDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetectionResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainDetectionResponse", string(data)}, " ")
}
