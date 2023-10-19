package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagKey 标签键，最大长度128个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，中文。
type TagKey struct {
}

func (o TagKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagKey struct{}"
	}

	return strings.Join([]string{"TagKey", string(data)}, " ")
}
