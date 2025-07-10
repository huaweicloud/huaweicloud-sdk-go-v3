package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScreenRecordsRequest Request Object
type ListScreenRecordsRequest struct {

	// 用于分页查询，返回录屏记录数量的限制。默认100。范围0~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 根据桌面ID过滤结果。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 根据用户名称过滤结果。
	Username *string `json:"username,omitempty"`

	// 录屏状态。 - RECORDING：录制中。 - REC_COMPLETED：录制完成。 - UPLOADING：上传中。 - UPLOAD_COMPLETED：上传完成。
	Status *string `json:"status,omitempty"`

	// 录屏类型。 - FULL：全程录屏。 - INTERVAL：间隔录屏。 - OPERATION：用户操作录屏。 - SESSION：监听会话生命周期录屏。
	Type *string `json:"type,omitempty"`

	// 开始时间，格式为yyyy-MM-dd HH:mm:ss（UTC时间，不传查默认最近15天）。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式为yyyy-MM-dd HH:mm:ss（UTC时间，不传查默认最近15天）。
	EndTime *string `json:"end_time,omitempty"`

	// 用于排序，表示按照哪个字段排序。取值为录屏属性start_time字段。
	SortField *string `json:"sort_field,omitempty"`

	// 用于排序，表示升序还是降序，取值为asc和desc。与sort_field一起组合使用。
	SortType *string `json:"sort_type,omitempty"`
}

func (o ListScreenRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScreenRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListScreenRecordsRequest", string(data)}, " ")
}
