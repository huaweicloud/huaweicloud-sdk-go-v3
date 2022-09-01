package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 待删除会场信息
type DelAttendInfo struct {

	// 会场呼叫号码。
	Number string `json:"number" xml:"number"`

	// 与会者标识，已入会的必须填写该字段。
	ParticipantID *string `json:"participantID,omitempty" xml:"participantID"`
}

func (o DelAttendInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelAttendInfo struct{}"
	}

	return strings.Join([]string{"DelAttendInfo", string(data)}, " ")
}
