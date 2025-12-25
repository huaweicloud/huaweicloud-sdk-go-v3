package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcDomain 云域名详情
type HwcDomain struct {

	// 域名名称
	DomainName string `json:"domain_name"`

	// 域名到期时间，eg：2023-01-10
	ExpireDate string `json:"expire_date"`

	// 域名服务状态
	Status string `json:"status"`

	// 域名实名认证状态。 取值范围： NONAUDIT：未实名认证 SUCCEED：已实名认证 FAILED：实名认证失败 AUDITING：实名认证审核中
	AuditStatus string `json:"audit_status"`

	// 域名实名认证失败原因
	AuditUnpassReason string `json:"audit_unpass_reason"`

	// 过户状态
	TransferStatus *string `json:"transfer_status,omitempty"`

	// 注册类型 取值范围： PERSONAL：个人 COMPANY：企业
	RegType string `json:"reg_type"`

	// 是否开启隐私保护
	PrivacyProtection string `json:"privacy_protection"`

	// 域名服务器列表
	NameServer []string `json:"name_server"`

	// 证件类型
	CredentialType string `json:"credential_type"`

	// 证件号码
	CredentialId string `json:"credential_id"`

	// 域名所属注册商
	Registrar string `json:"registrar"`

	// 联系人信息
	Contact []HwcDomainContact `json:"contact"`
}

func (o HwcDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcDomain struct{}"
	}

	return strings.Join([]string{"HwcDomain", string(data)}, " ")
}
