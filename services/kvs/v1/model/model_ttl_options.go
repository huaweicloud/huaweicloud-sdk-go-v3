package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TtlOptions struct {

	// TTL开关
	TtlSwitch bool `bson:"ttl_switch"`

	// 生存时间，以秒为单位
	ExpireAfterSeconds *int32 `bson:"expire_after_seconds,omitempty"`

	// 文档中记录TTL过期时间的字段名，字段值为UTC时间，单位秒
	TtlFieldName *string `bson:"ttl_field_name,omitempty"`
}

func (o TtlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtlOptions struct{}"
	}

	return strings.Join([]string{"TtlOptions", string(data)}, " ")
}
