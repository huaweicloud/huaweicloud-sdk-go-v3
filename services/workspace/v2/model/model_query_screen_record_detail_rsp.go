package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryScreenRecordDetailRsp struct {

	// 主键UUID。
	Id *string `json:"id,omitempty"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面池ID。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 用户名称。
	Username *string `json:"username,omitempty"`

	// 文件大小（Byte）。
	Size *int32 `json:"size,omitempty"`

	// 录屏类型。 - FULL：全程录屏。 - INTERVAL：间隔录屏。 - USER_OPERATION：用户操作录屏。 - SESSION：监听会话生命周期录屏。
	Type *string `json:"type,omitempty"`

	// 录屏状态。 - RECORDING：录制中。 - REC_COMPLETED：录制完成。 - REC_FAILED：录制失败。 - UPLOADING：上传中。 - UPLOAD_COMPLETED：上传完成。 - UPLOAD_FAILED：上传失败。
	Status *string `json:"status,omitempty"`

	// 录屏文件名称。
	VideoFilename *string `json:"video_filename,omitempty"`

	// 事件文件名称。
	EventFilename *string `json:"event_filename,omitempty"`

	// 开始时间（2024-10-15T10:04:41.263Z）。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间（2024-10-15T11:04:41.263Z）。
	EndTime *string `json:"end_time,omitempty"`

	// 更新时间（2024-10-15T11:04:41.263Z）。
	UpdateTime *string `json:"update_time,omitempty"`

	// 视频时长（秒）。
	Duration *int32 `json:"duration,omitempty"`
}

func (o QueryScreenRecordDetailRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryScreenRecordDetailRsp struct{}"
	}

	return strings.Join([]string{"QueryScreenRecordDetailRsp", string(data)}, " ")
}
