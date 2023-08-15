package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetRect 目标检测框信息。
type TargetRect struct {

	// **参数说明**：摄像头编码。
	CameraCode *string `json:"camera_code,omitempty"`

	// **参数说明**：摄像头IP地址。
	CameraIp *string `json:"camera_ip,omitempty"`

	TargetPos *TargetPos `json:"target_pos,omitempty"`

	// **参数说明**：与SnapTime的时间差值：当前检测框所在相机的时间戳减去雷视拟合轨迹中的SnapTime的差值。
	TimeStampDiff *int64 `json:"time_stamp_diff,omitempty"`
}

func (o TargetRect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetRect struct{}"
	}

	return strings.Join([]string{"TargetRect", string(data)}, " ")
}
