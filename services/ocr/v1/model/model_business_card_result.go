package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BusinessCardResult struct {

	// 姓名列表。
	Name *[]string `json:"name,omitempty" xml:"name"`

	// 职位头衔列表。
	Title *[]string `json:"title,omitempty" xml:"title"`

	// 公司列表。
	Company *[]string `json:"company,omitempty" xml:"company"`

	// 部门列表。
	Department *[]string `json:"department,omitempty" xml:"department"`

	// 联系方式列表。
	Phone *[]string `json:"phone,omitempty" xml:"phone"`

	// 地址列表。
	Address *[]string `json:"address,omitempty" xml:"address"`

	// 邮箱列表。
	Email *[]string `json:"email,omitempty" xml:"email"`

	// 传真列表。
	Fax *[]string `json:"fax,omitempty" xml:"fax"`

	// 邮编列表。
	Postcode *[]string `json:"postcode,omitempty" xml:"postcode"`

	// 公司网址列表。
	Website *[]string `json:"website,omitempty" xml:"website"`

	// 其余信息列表。
	ExtraInfoList *[]ExtraInfoList `json:"extra_info_list,omitempty" xml:"extra_info_list"`

	// 返回矫正后的名片图像的BASE64编码。
	AdjustedImage *string `json:"adjusted_image,omitempty" xml:"adjusted_image"`
}

func (o BusinessCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCardResult struct{}"
	}

	return strings.Join([]string{"BusinessCardResult", string(data)}, " ")
}
