package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessControlRule struct {

	// IP地址或网段。
	Ip string `json:"ip"`

	// 备注。备注长度范围是0到50个字符，不能包含<>。
	Description *string `json:"description,omitempty"`
}

func (o AccessControlRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControlRule struct{}"
	}

	return strings.Join([]string{"AccessControlRule", string(data)}, " ")
}
