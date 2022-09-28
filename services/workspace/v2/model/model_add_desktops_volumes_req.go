package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 增加磁盘请求。
type AddDesktopsVolumesReq struct {

	// 新增磁盘参数。
	AddDesktopVolumesReq *[]AddDesktopVolumesReq `json:"addDesktopVolumesReq,omitempty"`
}

func (o AddDesktopsVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopsVolumesReq struct{}"
	}

	return strings.Join([]string{"AddDesktopsVolumesReq", string(data)}, " ")
}
