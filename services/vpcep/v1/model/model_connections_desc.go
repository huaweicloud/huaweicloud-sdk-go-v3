package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionsDesc struct {

	// 终端节点ID，UUID格式字符
	Id string `json:"id"`

	// 描述字段，支持中英文字母、数字等字符，不支持“<”或“>”字符。
	Description string `json:"description"`
}

func (o ConnectionsDesc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionsDesc struct{}"
	}

	return strings.Join([]string{"ConnectionsDesc", string(data)}, " ")
}
