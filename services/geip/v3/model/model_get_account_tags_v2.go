package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAccountTagsV2 资源标签。
type GetAccountTagsV2 struct {

	// 标签键，最大长度128个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，中文。
	Key string `json:"key"`

	// tag的value列表
	Values []string `json:"values"`
}

func (o GetAccountTagsV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccountTagsV2 struct{}"
	}

	return strings.Join([]string{"GetAccountTagsV2", string(data)}, " ")
}
