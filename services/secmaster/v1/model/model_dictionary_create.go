package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DictionaryCreate struct {

	// 版本号
	Version *string `json:"version,omitempty"`

	// 字典id
	DictId string `json:"dict_id"`

	// 字典key
	DictKey string `json:"dict_key"`

	// 字典code码
	DictCode string `json:"dict_code"`

	// 字典值
	DictVal string `json:"dict_val"`

	// 字典key
	DictPkey *string `json:"dict_pkey,omitempty"`

	// 字典关联的父code
	DictPcode *string `json:"dict_pcode,omitempty"`

	// 字典所属领域
	Scope *string `json:"scope,omitempty"`

	// 字典描述信息
	Description *string `json:"description,omitempty"`

	// 额外字段
	ExtendField *interface{} `json:"extend_field,omitempty"`

	// 用户当前语言环境
	Language string `json:"language"`
}

func (o DictionaryCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryCreate struct{}"
	}

	return strings.Join([]string{"DictionaryCreate", string(data)}, " ")
}
