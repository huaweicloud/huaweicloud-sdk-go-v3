package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IntellectualPropertyV2 struct {

	// 知识产权类型,分：著作权-copyright、商标权-trademark、专利权-patent
	IpType *string `json:"ip_type,omitempty"`

	// 知识产权注册国家/地区
	IpArea *string `json:"ip_area,omitempty"`

	// 商标注册号/专利申请号
	IpNumber *string `json:"ip_number,omitempty"`

	// 知识产权情况简述
	IpContent *string `json:"ip_content,omitempty"`
}

func (o IntellectualPropertyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntellectualPropertyV2 struct{}"
	}

	return strings.Join([]string{"IntellectualPropertyV2", string(data)}, " ")
}
