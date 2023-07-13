package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsedDesktopInfoReq 查询使用桌面时长请求。
type ListUsedDesktopInfoReq struct {

	// 桌面id集合。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 开始时间，格式：yyyy-MM-dd（UTC时间，不传查默认最近15天）最多查31天数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式：yyyy-MM-dd（UTC时间，不传查默认最近15天）最多查31天数据。
	EndTime *string `json:"end_time,omitempty"`

	// 若传桌面的用户名，则查询使用时间只有该用户的使用时间。
	DesktopUsername *string `json:"desktop_username,omitempty"`

	// 从查询结果中的第几条数据开始返回,用于分页查询，取值范围0-2000，默认从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果中想要返回的信息条目数量,用于分页查询，取值范围0-100，默认值100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUsedDesktopInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsedDesktopInfoReq struct{}"
	}

	return strings.Join([]string{"ListUsedDesktopInfoReq", string(data)}, " ")
}
