package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolVolumesReqVolumes 磁盘信息。
type ExpandDesktopPoolVolumesReqVolumes struct {

	// 批量操作磁盘的磁盘集合id。
	Id string `json:"id"`

	// 磁盘容量，单位GB。
	Size int32 `json:"size"`
}

func (o ExpandDesktopPoolVolumesReqVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolVolumesReqVolumes struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolVolumesReqVolumes", string(data)}, " ")
}
