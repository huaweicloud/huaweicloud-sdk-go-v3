package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOrgResResponse struct {

	// 企业管理员查询所属企业的资源使用信息。
	UsedRecStorage *float64 `json:"usedRecStorage,omitempty"`

	// 当前已用的会议并发数量。
	UsedAccountsCount *int32 `json:"usedAccountsCount,omitempty"`

	// 当前已用的直播推流资源。
	UsedLiveCount *int32 `json:"usedLiveCount,omitempty"`

	// 当前已用的直播推流资源。
	ConfCount *int32 `json:"confCount,omitempty"`

	// 当日会议总时长。
	ConfLength *int64 `json:"confLength,omitempty"`

	// 当日活跃用户数。
	ActiveAttendeeCount *int32 `json:"activeAttendeeCount,omitempty"`

	// 当日总与会人数。
	TotalAttendeeCount *int32 `json:"totalAttendeeCount,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o ShowOrgResResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrgResResponse struct{}"
	}

	return strings.Join([]string{"ShowOrgResResponse", string(data)}, " ")
}
