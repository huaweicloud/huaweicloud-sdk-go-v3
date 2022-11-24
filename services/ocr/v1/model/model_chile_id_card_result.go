package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChileIdCardResult struct {

	// 姓氏。
	Surname *[]string `json:"surname,omitempty"`

	// 名。
	GivenName *string `json:"given_name,omitempty"`

	// 国籍。
	Nationality *string `json:"nationality,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 出生日。
	Birth *string `json:"birth,omitempty"`

	// 发行日。
	IssueDate *string `json:"issue_date,omitempty"`

	// 有效期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 文档编号。
	DocumentNumber *string `json:"document_number,omitempty"`

	// 身份证号。
	Number *string `json:"number,omitempty"`

	Confidence *ChileIdCardConfidence `json:"confidence,omitempty"`
}

func (o ChileIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChileIdCardResult struct{}"
	}

	return strings.Join([]string{"ChileIdCardResult", string(data)}, " ")
}
