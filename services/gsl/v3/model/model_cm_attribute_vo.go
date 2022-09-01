package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CmAttributeVo struct {

	// 自定义属性标识
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 默认属性名称中文
	DefaultAttrNameCn *string `json:"default_attr_name_cn,omitempty" xml:"default_attr_name_cn"`

	// 默认属性名称英文
	DefaultAttrNameEn *string `json:"default_attr_name_en,omitempty" xml:"default_attr_name_en"`

	// 自定义属性名称
	CustAttrName *string `json:"cust_attr_name,omitempty" xml:"cust_attr_name"`

	// 自定义属性状态：0 未启用，1 已启用。
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间
	ModifyTime *sdktime.SdkTime `json:"modify_time,omitempty" xml:"modify_time"`
}

func (o CmAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CmAttributeVo struct{}"
	}

	return strings.Join([]string{"CmAttributeVo", string(data)}, " ")
}
