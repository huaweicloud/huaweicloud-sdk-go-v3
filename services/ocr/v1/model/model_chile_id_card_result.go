package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChileIdCardResult struct {

	// 姓氏。
	Surname *[]string `json:"surname,omitempty" xml:"surname"`

	// 名。
	GivenName *string `json:"given_name,omitempty" xml:"given_name"`

	// 国籍。
	Nationality *string `json:"nationality,omitempty" xml:"nationality"`

	// 性别。
	Sex *string `json:"sex,omitempty" xml:"sex"`

	// 出生日。
	Birth *string `json:"birth,omitempty" xml:"birth"`

	// 发行日。
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date"`

	// 有效期。
	ExpiryDate *string `json:"expiry_date,omitempty" xml:"expiry_date"`

	// 文档编号。
	DocumentNumber *string `json:"document_number,omitempty" xml:"document_number"`

	// 身份证号。
	Number *string `json:"number,omitempty" xml:"number"`

	Confidence *ChileIdCardConfidence `json:"confidence,omitempty" xml:"confidence"`
}

func (o ChileIdCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChileIdCardResult struct{}"
	}

	return strings.Join([]string{"ChileIdCardResult", string(data)}, " ")
}
