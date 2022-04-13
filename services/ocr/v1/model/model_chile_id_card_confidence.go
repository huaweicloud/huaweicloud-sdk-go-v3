package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChileIdCardConfidence struct {
	// 姓氏置信度。

	Surname *float32 `json:"surname,omitempty"`
	// 名置信度。

	GivenName *float32 `json:"given_name,omitempty"`
	// 国籍置信度。

	Nationality *float32 `json:"nationality,omitempty"`
	// 性别置信度。

	Sex *float32 `json:"sex,omitempty"`
	// 出生日置信度。

	Birth *float32 `json:"birth,omitempty"`
	// 发行日置信度。

	IssueDate *float32 `json:"issue_date,omitempty"`
	// 有效期置信度。

	ExpiryDate *float32 `json:"expiry_date,omitempty"`
	// 文档编号置信度。

	DocumentNumber *float32 `json:"document_number,omitempty"`
	// 身份证号置信度。

	Number *float32 `json:"number,omitempty"`
}

func (o ChileIdCardConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChileIdCardConfidence struct{}"
	}

	return strings.Join([]string{"ChileIdCardConfidence", string(data)}, " ")
}
