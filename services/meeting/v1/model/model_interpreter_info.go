package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InterpreterInfo 传译员信息
type InterpreterInfo struct {

	// 用户登录账号，可以是账号、手机、邮箱其中一个,loginAccount和userID必须二选一。
	LoginAccount string `json:"loginAccount"`

	// 用户的userUUID。
	UserID *string `json:"userID,omitempty"`

	// 呼叫号码。
	CallNumber *string `json:"callNumber,omitempty"`

	// 用户名。
	Name *string `json:"name,omitempty"`

	// 传译员序号，AI传译组下传译员唯一标识。
	InterpreterId *string `json:"interpreterId,omitempty"`
}

func (o InterpreterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterpreterInfo struct{}"
	}

	return strings.Join([]string{"InterpreterInfo", string(data)}, " ")
}
