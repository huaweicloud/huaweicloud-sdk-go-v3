package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOrgResResponse struct {

	// 已用录制存储空间(单位G)
	UsedRecStorage *float64 `json:"usedRecStorage,omitempty" xml:"usedRecStorage"`

	// 已用的企业并发数
	UsedAccountsCount *int32 `json:"usedAccountsCount,omitempty" xml:"usedAccountsCount"`

	// 已用的直播推流资源
	UsedLiveCount *int32 `json:"usedLiveCount,omitempty" xml:"usedLiveCount"`

	// 会议总次数
	ConfCount *int32 `json:"confCount,omitempty" xml:"confCount"`

	// 会议总时长
	ConfLength *int64 `json:"confLength,omitempty" xml:"confLength"`

	// 活跃用户数
	ActiveAttendeeCount *int32 `json:"activeAttendeeCount,omitempty" xml:"activeAttendeeCount"`

	// 总与会人数
	TotalAttendeeCount *int32 `json:"totalAttendeeCount,omitempty" xml:"totalAttendeeCount"`
	HttpStatusCode     int    `json:"-"`
}

func (o ShowOrgResResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrgResResponse struct{}"
	}

	return strings.Join([]string{"ShowOrgResResponse", string(data)}, " ")
}
