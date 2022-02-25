package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量删除member请求参数。
type BatchDeleteMembersOption struct {
	// 需要删除的后端服务器ID。  使用说明： - 若传入id则不能传其他参数，否则报错。

	Id *string `json:"id,omitempty"`
	// 后端服务器IP地址。  使用说明： - address和protocol_port必须同时传入。 - 不能同时传入ID字段

	Address *string `json:"address,omitempty"`
	// 后端服务器端口。  使用说明： - address和protocol_port必须同时传入。 - 不能同时传入ID字段

	ProtocolPort *int32 `json:"protocol_port,omitempty"`
}

func (o BatchDeleteMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersOption", string(data)}, " ")
}
