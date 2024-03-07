package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagValue 与标签键关联的字符串值。
type TagValue struct {
}

func (o TagValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValue struct{}"
	}

	return strings.Join([]string{"TagValue", string(data)}, " ")
}
