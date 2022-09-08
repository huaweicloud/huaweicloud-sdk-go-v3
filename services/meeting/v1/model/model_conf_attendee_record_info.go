package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与会者记录
type ConfAttendeeRecordInfo struct {

	// 名称。
	DisplayName *string `json:"displayName,omitempty"`

	// 号码。
	CallNumber *string `json:"callNumber,omitempty"`

	// 设备类型。
	DeviceType *string `json:"deviceType,omitempty"`

	// 入会时间（UTC时间，单位毫秒）。
	JoinTime *int64 `json:"joinTime,omitempty"`

	// 离会时间（UTC时间，单位毫秒）。
	LeftTime *int64 `json:"leftTime,omitempty"`

	// 媒体类型。
	MediaType *string `json:"mediaType,omitempty"`

	// 部门名称。
	DeptName *string `json:"deptName,omitempty"`
}

func (o ConfAttendeeRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfAttendeeRecordInfo struct{}"
	}

	return strings.Join([]string{"ConfAttendeeRecordInfo", string(data)}, " ")
}
