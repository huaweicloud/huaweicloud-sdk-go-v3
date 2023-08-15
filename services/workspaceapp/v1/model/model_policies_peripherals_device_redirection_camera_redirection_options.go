package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionCameraRedirectionOptions 摄像头设备重定向控制的选项。
type PoliciesPeripheralsDeviceRedirectionCameraRedirectionOptions struct {

	// 摄像头帧率（fps）。取值范围为[1-30]。默认：15。
	CameraFrameRate *int32 `json:"camera_frame_rate,omitempty"`

	// 摄像头最大宽度（pixel）。取值范围为[1-9999]。默认：3000。
	CameraMaxWidth *int32 `json:"camera_max_width,omitempty"`

	// 摄像头最大高度（pixel）。取值范围为[1-9999]。默认：3000。
	CameraMaxHeigth *int32 `json:"camera_max_heigth,omitempty"`

	// 摄像头数据压缩方式。取值为：H.264。
	CameraCompressionMethod *string `json:"camera_compression_method,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionCameraRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionCameraRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionCameraRedirectionOptions", string(data)}, " ")
}
