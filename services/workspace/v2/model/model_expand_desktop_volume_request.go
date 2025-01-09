package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopVolumeRequest Request Object
type ExpandDesktopVolumeRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	// 磁盘ID。
	VolumeId string `json:"volume_id"`

	Body *ExpandVolumeReq `json:"body,omitempty"`
}

func (o ExpandDesktopVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopVolumeRequest struct{}"
	}

	return strings.Join([]string{"ExpandDesktopVolumeRequest", string(data)}, " ")
}
