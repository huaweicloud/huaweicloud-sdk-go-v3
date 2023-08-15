package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DelAttendInfo 待删除会场信息
type DelAttendInfo struct {

	// 会场号码。
	Number string `json:"number"`

	// 与会者标识，已入会的必须填写该字段。
	ParticipantID *string `json:"participantID,omitempty"`
}

func (o DelAttendInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelAttendInfo struct{}"
	}

	return strings.Join([]string{"DelAttendInfo", string(data)}, " ")
}
