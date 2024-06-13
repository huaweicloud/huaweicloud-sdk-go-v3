package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetPublicLibAndAwsResp struct {

	// 公共关键字分组信息
	AwTag *string `json:"aw_tag,omitempty"`

	// 公共关键字帮助文档链接
	DocumentLink *string `json:"document_link,omitempty"`

	// 保留字段
	IsFavorite *int32 `json:"is_favorite,omitempty"`

	// 公共关键字描述
	PublicAwDescription *string `json:"public_aw_description,omitempty"`

	// 公共关键字唯一ID
	PublicAwId *string `json:"public_aw_id,omitempty"`

	// 公共关键库唯一ID
	PublicAwLibId *string `json:"public_aw_lib_id,omitempty"`

	// 公共关键字库名称
	PublicAwLibName *string `json:"public_aw_lib_name,omitempty"`

	// 保留字段
	PublicAwMark *int32 `json:"public_aw_mark,omitempty"`

	// 公共关键字名称
	PublicAwName *string `json:"public_aw_name,omitempty"`
}

func (o GetPublicLibAndAwsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPublicLibAndAwsResp struct{}"
	}

	return strings.Join([]string{"GetPublicLibAndAwsResp", string(data)}, " ")
}
