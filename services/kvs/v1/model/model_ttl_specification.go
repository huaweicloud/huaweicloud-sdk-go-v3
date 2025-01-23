package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtlSpecification TTL设置。非必填项，默认不启用TTL。
type TtlSpecification struct {

	// TTL开关
	Enable bool `bson:"enable"`

	// 生存时间，以秒为单位
	ExpireAfterSeconds *int32 `bson:"expire_after_seconds,omitempty"`

	// 文档中记录TTL过期时间的字段名，字段值为UTC时间，单位秒
	FieldName *string `bson:"field_name,omitempty"`
}

func (o TtlSpecification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtlSpecification struct{}"
	}

	return strings.Join([]string{"TtlSpecification", string(data)}, " ")
}
