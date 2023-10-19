package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagValue 标签值，最大长度255个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，点“.”，中文。
type TagValue struct {
}

func (o TagValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValue struct{}"
	}

	return strings.Join([]string{"TagValue", string(data)}, " ")
}
