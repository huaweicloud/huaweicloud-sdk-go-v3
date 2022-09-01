package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端状态信息
type UserStatusDto struct {

	// 终端号码
	Number *string `json:"Number,omitempty" xml:"Number"`

	// 注册状态。 * 1是未注册上 * 0是已注册
	RegStatus *string `json:"RegStatus,omitempty" xml:"RegStatus"`

	// 呼叫状态。 * 0:未上线 * 1:空闲中 * 2:使用中 * 3:非会议硬终端统一的无效值
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus"`
}

func (o UserStatusDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserStatusDto struct{}"
	}

	return strings.Join([]string{"UserStatusDto", string(data)}, " ")
}
