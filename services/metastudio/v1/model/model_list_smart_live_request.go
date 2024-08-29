package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartLiveRequest Request Object
type ListSmartLiveRequest struct {

	// 剧本ID。
	RoomId string `json:"room_id"`

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

	// 任务状态，默认所有状态。  可多个状态查询，使用英文逗号分隔。  如state=CREATING,PUBLISHED
	State *string `json:"state,omitempty"`

	// 排序字段，支持的排序方式有： - 按创建时间排序：create_time - 按更新时间排序：update_time - 按资产排序：asset_order
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式。 * asc：升序 * desc：降序  默认asc升序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 过滤创建时间>=输入时间的记录。
	CreateSince *string `json:"create_since,omitempty"`

	// 过滤创建时间<=输入时间的记录。
	CreateUntil *string `json:"create_until,omitempty"`
}

func (o ListSmartLiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartLiveRequest struct{}"
	}

	return strings.Join([]string{"ListSmartLiveRequest", string(data)}, " ")
}
