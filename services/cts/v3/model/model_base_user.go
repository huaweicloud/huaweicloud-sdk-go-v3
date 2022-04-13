package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户信息。
type BaseUser struct {
	// 账号ID，参见《云审计服务API参考》“获取账号ID和项目ID”章节。

	Id *string `json:"id,omitempty"`
	// 账号名称。

	Name *string `json:"name,omitempty"`
}

func (o BaseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseUser struct{}"
	}

	return strings.Join([]string{"BaseUser", string(data)}, " ")
}
