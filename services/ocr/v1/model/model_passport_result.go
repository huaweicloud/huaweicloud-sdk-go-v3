package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PassportResult struct {

	// 护照类型（P:普通因私护照、W:外交护照、G:公务护照）（英文）。
	PassportType *string `json:"passport_type,omitempty" xml:"passport_type"`

	// 护照签发国的国家码（英文）。
	CountryCode *string `json:"country_code,omitempty" xml:"country_code"`

	// 护照号码（英文）。
	PassportNumber *string `json:"passport_number,omitempty" xml:"passport_number"`

	// 护照持有人国籍（英文）。
	Nationality *string `json:"nationality,omitempty" xml:"nationality"`

	// 姓（英文）。
	Surname *string `json:"surname,omitempty" xml:"surname"`

	// 名字（英文）。
	GivenName *string `json:"given_name,omitempty" xml:"given_name"`

	// 性别（英文）。
	Sex *string `json:"sex,omitempty" xml:"sex"`

	// 出生日期（英文）。
	DateOfBirth *string `json:"date_of_birth,omitempty" xml:"date_of_birth"`

	// 护照有效期（英文）。
	DateOfExpiry *string `json:"date_of_expiry,omitempty" xml:"date_of_expiry"`

	// 护照签发日期（英文）。
	DateOfIssue *string `json:"date_of_issue,omitempty" xml:"date_of_issue"`

	// 出生地（英文）。
	PlaceOfBirth *string `json:"place_of_birth,omitempty" xml:"place_of_birth"`

	// 签发地（英文）。
	PlaceOfIssue *string `json:"place_of_issue,omitempty" xml:"place_of_issue"`

	// 签发机构（英文），其中对中国的英文简写统一输出为P.R.China。
	IssuingAuthority *string `json:"issuing_authority,omitempty" xml:"issuing_authority"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty" xml:"confidence"`

	// 默认为空。对于部分常见国家的护照OCR服务，extra_info内会包含护照上由本地官方语言描述的字段信息及其他信息。 如中国护照，里面会包含汉字表达的姓名、出生地等信息。
	ExtraInfo *interface{} `json:"extra_info,omitempty" xml:"extra_info"`
}

func (o PassportResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PassportResult struct{}"
	}

	return strings.Join([]string{"PassportResult", string(data)}, " ")
}
