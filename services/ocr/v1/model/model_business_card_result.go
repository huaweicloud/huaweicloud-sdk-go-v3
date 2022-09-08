package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BusinessCardResult struct {

	// 姓名列表。
	Name *[]string `json:"name,omitempty"`

	// 职位头衔列表。
	Title *[]string `json:"title,omitempty"`

	// 公司列表。
	Company *[]string `json:"company,omitempty"`

	// 部门列表。
	Department *[]string `json:"department,omitempty"`

	// 联系方式列表。
	Phone *[]string `json:"phone,omitempty"`

	// 地址列表。
	Address *[]string `json:"address,omitempty"`

	// 邮箱列表。
	Email *[]string `json:"email,omitempty"`

	// 传真列表。
	Fax *[]string `json:"fax,omitempty"`

	// 邮编列表。
	Postcode *[]string `json:"postcode,omitempty"`

	// 公司网址列表。
	Website *[]string `json:"website,omitempty"`

	// 其余信息列表。
	ExtraInfoList *[]ExtraInfoList `json:"extra_info_list,omitempty"`

	// 返回矫正后的名片图像的BASE64编码。
	AdjustedImage *string `json:"adjusted_image,omitempty"`
}

func (o BusinessCardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCardResult struct{}"
	}

	return strings.Join([]string{"BusinessCardResult", string(data)}, " ")
}
