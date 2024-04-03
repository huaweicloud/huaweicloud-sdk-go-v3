package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Message struct {

	// 角色
	Role *string `json:"role,omitempty"`

	// 问答对文本内容，最小长度：1，最大长度：模型支持的max_tokens数量乘以系数，默认系数为1.5，并且所有content的总长度不能超过该最大长度
	Content string `json:"content"`
}

func (o Message) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Message struct{}"
	}

	return strings.Join([]string{"Message", string(data)}, " ")
}
