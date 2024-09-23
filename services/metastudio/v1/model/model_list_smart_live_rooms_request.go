package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartLiveRoomsRequest Request Object
type ListSmartLiveRoomsRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 按直播间名称模糊查询。
	RoomName *string `json:"room_name,omitempty"`

	// 按数字人形象ID查询。
	DhId *string `json:"dh_id,omitempty"`

	// 按形象名称模糊查询。
	ModelName *string `json:"model_name,omitempty"`

	// 当前直播间直播状态。 WAITING，PROCESSING，SUCCESS，FAILED，CANCELED对应直播任务状态 NULL 对应没有直播任务 可多个状态查询，使用英文逗号分隔。
	LiveState *string `json:"live_state,omitempty"`

	// 最近直播任务起始时间。格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。格式遵循：RFC 3339 如\"2021-01-10T10:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	// 按直播间类型查询。直播间类型。 * NORMAL：普通直播间，直播间一直存在，可以反复开播 * TEMP：临时直播间，直播任务结束后自动清理直播间。 * TEMPLATE：直播间模板。
	RoomType *string `json:"room_type,omitempty"`

	// 按照自己拥有的和别人分享以及公共的模板进行查询 * OWNED 自己拥有且暂未共享的 * SHARED_TO_OHTERS 分享给别人的 * SHARED_FROM_OHTERS 别人分享给我的 * PUBLIC 公共模板
	TemplateOwnType *string `json:"template_own_type,omitempty"`
}

func (o ListSmartLiveRoomsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartLiveRoomsRequest struct{}"
	}

	return strings.Join([]string{"ListSmartLiveRoomsRequest", string(data)}, " ")
}
