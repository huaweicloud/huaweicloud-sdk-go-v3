package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChileIdCardConfidence struct {

	// 姓氏置信度。
	Surname *float32 `json:"surname,omitempty" xml:"surname"`

	// 名置信度。
	GivenName *float32 `json:"given_name,omitempty" xml:"given_name"`

	// 国籍置信度。
	Nationality *float32 `json:"nationality,omitempty" xml:"nationality"`

	// 性别置信度。
	Sex *float32 `json:"sex,omitempty" xml:"sex"`

	// 出生日置信度。
	Birth *float32 `json:"birth,omitempty" xml:"birth"`

	// 发行日置信度。
	IssueDate *float32 `json:"issue_date,omitempty" xml:"issue_date"`

	// 有效期置信度。
	ExpiryDate *float32 `json:"expiry_date,omitempty" xml:"expiry_date"`

	// 文档编号置信度。
	DocumentNumber *float32 `json:"document_number,omitempty" xml:"document_number"`

	// 身份证号置信度。
	Number *float32 `json:"number,omitempty" xml:"number"`
}

func (o ChileIdCardConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChileIdCardConfidence struct{}"
	}

	return strings.Join([]string{"ChileIdCardConfidence", string(data)}, " ")
}
