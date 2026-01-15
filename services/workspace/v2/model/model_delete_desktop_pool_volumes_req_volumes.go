package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopPoolVolumesReqVolumes 磁盘信息。
type DeleteDesktopPoolVolumesReqVolumes struct {

	// 批量操作磁盘的磁盘集合id。
	Id string `json:"id"`
}

func (o DeleteDesktopPoolVolumesReqVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopPoolVolumesReqVolumes struct{}"
	}

	return strings.Join([]string{"DeleteDesktopPoolVolumesReqVolumes", string(data)}, " ")
}
