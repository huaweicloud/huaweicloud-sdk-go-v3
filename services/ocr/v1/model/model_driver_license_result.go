package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DriverLicenseResult struct {
	// 驾驶证号。

	Number *string `json:"number,omitempty"`
	// 姓名。

	Name *string `json:"name,omitempty"`
	// 性别。

	Sex *string `json:"sex,omitempty"`
	// 国籍。

	Nationality *string `json:"nationality,omitempty"`
	// 住址。

	Address *string `json:"address,omitempty"`
	// 出生日期。

	Birth *string `json:"birth,omitempty"`
	// 初次领证日期。

	IssueDate *string `json:"issue_date,omitempty"`
	// 准驾类型。

	Class *string `json:"class,omitempty"`
	// 有效起始日期。

	ValidFrom *string `json:"valid_from,omitempty"`
	// 有效结束日期。

	ValidTo *string `json:"valid_to,omitempty"`
	// 发证机关。

	IssuingAuthority *string `json:"issuing_authority,omitempty"`
	// 档案编号。

	FileNumber *string `json:"file_number,omitempty"`
	// 记录。

	Record *string `json:"record,omitempty"`
}

func (o DriverLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriverLicenseResult struct{}"
	}

	return strings.Join([]string{"DriverLicenseResult", string(data)}, " ")
}
