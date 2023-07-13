package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionCameraRedirection 摄像头重定向。
type PoliciesPeripheralsDeviceRedirectionCameraRedirection struct {

	// 是否开启摄像头设备重定向。取值为： false：表示关闭。 true：表示开启。
	VideoCompressEnable *bool `json:"video_compress_enable,omitempty"`

	Options *PoliciesPeripheralsDeviceRedirectionCameraRedirectionOptions `json:"options,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionCameraRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionCameraRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionCameraRedirection", string(data)}, " ")
}
