package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateResponseBodyApplicationInfo 证书申请信息。
type ShowCertificateResponseBodyApplicationInfo struct {

	// 国家。
	Country *string `json:"country,omitempty"`

	// 公司名称。
	CompanyName *string `json:"company_name,omitempty"`

	// 公司省份。
	CompanyProvince *string `json:"company_province,omitempty"`

	// 公司所在城市。
	CompanyCity *string `json:"company_city,omitempty"`

	// 申请人名称。
	ApplicantName *string `json:"applicant_name,omitempty"`

	// 申请人电话。
	ApplicantPhone *string `json:"applicant_phone,omitempty"`

	// 申请人邮箱。
	ApplicantEmail *string `json:"applicant_email,omitempty"`

	// 技术联系人名称。
	ContactName *string `json:"contact_name,omitempty"`

	// 技术联系人电话。
	ContactPhone *string `json:"contact_phone,omitempty"`

	// 技术联系人邮箱。
	ContactEmail *string `json:"contact_email,omitempty"`
}

func (o ShowCertificateResponseBodyApplicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateResponseBodyApplicationInfo struct{}"
	}

	return strings.Join([]string{"ShowCertificateResponseBodyApplicationInfo", string(data)}, " ")
}
