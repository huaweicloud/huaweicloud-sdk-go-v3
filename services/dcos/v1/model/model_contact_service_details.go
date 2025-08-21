package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContactServiceDetails 人员联系信息
type ContactServiceDetails struct {

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

	// 状态
	TaskStatus *string `json:"task_status,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o ContactServiceDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContactServiceDetails struct{}"
	}

	return strings.Join([]string{"ContactServiceDetails", string(data)}, " ")
}
