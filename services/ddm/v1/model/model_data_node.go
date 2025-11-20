package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataNode struct {

	// 实例id。
	Id *string `json:"id,omitempty"`

	// 实例账号。
	User *string `json:"user,omitempty"`

	// 实例账号密码。
	Password *string `json:"password,omitempty"`
}

func (o DataNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataNode struct{}"
	}

	return strings.Join([]string{"DataNode", string(data)}, " ")
}
