package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopsVolumesReq 增加磁盘请求。
type AddDesktopsVolumesReq struct {

	// 新增磁盘参数。
	DesktopVolumes *[]AddDesktopVolumesReq `json:"desktop_volumes,omitempty"`
}

func (o AddDesktopsVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopsVolumesReq struct{}"
	}

	return strings.Join([]string{"AddDesktopsVolumesReq", string(data)}, " ")
}
