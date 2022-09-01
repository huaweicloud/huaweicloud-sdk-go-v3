package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckDictionaryResponse struct {

	// 字典ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 字典名称 - 字符集：中文、英文字母、数字、下划线和空格 - 约束：实例下唯一
	Name *string `json:"name,omitempty" xml:"name"`

	// 字典描述
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 字典编码 - 字符集：英文字母、数字、下划线和空格 - 约束：实例下唯一
	Code *string `json:"code,omitempty" xml:"code"`

	// 字典扩展字段1 - 字符集：中文、英文字母、数字、下划线和空格
	ExtendOne *string `json:"extend_one,omitempty" xml:"extend_one"`

	// 字典扩展字段2 - 字符集：中文、英文字母、数字、下划线和空格
	ExtendTwo *string `json:"extend_two,omitempty" xml:"extend_two"`

	// 父字典编码,为空时代表自身就是最顶级字典
	ParentCode *string `json:"parent_code,omitempty" xml:"parent_code"`

	Type *DictionaryType `json:"type,omitempty" xml:"type"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckDictionaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDictionaryResponse struct{}"
	}

	return strings.Join([]string{"CheckDictionaryResponse", string(data)}, " ")
}
