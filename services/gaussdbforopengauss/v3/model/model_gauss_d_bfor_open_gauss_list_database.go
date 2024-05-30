package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussListDatabase 数据库信息。
type GaussDBforOpenGaussListDatabase struct {

	// 数据库名称。
	Name *string `json:"name,omitempty"`

	// 数据库所属用户。
	Owner *string `json:"owner,omitempty"`

	// 数据库使用的字符集，例如UTF8。
	CharacterSet *string `json:"character_set,omitempty"`

	// 数据库排序集，例如en_US.UTF-8等。
	CollateSet *string `json:"collate_set,omitempty"`

	// 数据库大小（单位：MB）。
	Size *string `json:"size,omitempty"`

	// 数据库使用的字符分类，例如en_US.UTF-8等。
	Datctype *string `json:"datctype,omitempty"`

	// 数据库兼容的类型，如GaussDB，M。
	CompatibilityType *string `json:"compatibility_type,omitempty"`
}

func (o GaussDBforOpenGaussListDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussListDatabase struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussListDatabase", string(data)}, " ")
}
