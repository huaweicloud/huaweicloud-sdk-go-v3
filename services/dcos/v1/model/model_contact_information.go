package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContactInformation 人员联系信息
type ContactInformation struct {

	// 人员姓名
	Name string `json:"name"`

	// 证件号前三位
	CertificateNoTop *string `json:"certificate_no_top,omitempty"`

	// 证件号后四位
	CertificateNoLast *string `json:"certificate_no_last,omitempty"`

	// 手机国际码
	CountryCode *string `json:"country_code,omitempty"`

	// 电话
	Phone *string `json:"phone,omitempty"`

	// 单位
	WorkUnit *string `json:"work_unit,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`
}

func (o ContactInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContactInformation struct{}"
	}

	return strings.Join([]string{"ContactInformation", string(data)}, " ")
}
