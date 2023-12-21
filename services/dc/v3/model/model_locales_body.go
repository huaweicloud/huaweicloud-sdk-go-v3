package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocalesBody Locale地址描述信息。
type LocalesBody struct {

	// 区域英文名称。
	EnUs *string `json:"en_us,omitempty"`

	// 区域中文名称。
	ZhCn *string `json:"zh_cn,omitempty"`
}

func (o LocalesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalesBody struct{}"
	}

	return strings.Join([]string{"LocalesBody", string(data)}, " ")
}
