package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 高级配置
type Advance struct {
	// 索引（参数：params，会话cookie：cookie，header字段：header，body字段：body，多种组合：multipart）

	Index *string `json:"index,omitempty"`
	// 指定字段（仅在param，cookie，header模式下可以使用）

	Contents *[]string `json:"contents,omitempty"`
}

func (o Advance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Advance struct{}"
	}

	return strings.Join([]string{"Advance", string(data)}, " ")
}
