package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 扩容磁盘请求。
type ExpandDesktopsVolumesReq struct {

	// 扩容磁盘参数。
	ExpandVolumesReq *[]ExpandVolumesReq `json:"expandVolumesReq,omitempty"`
}

func (o ExpandDesktopsVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopsVolumesReq struct{}"
	}

	return strings.Join([]string{"ExpandDesktopsVolumesReq", string(data)}, " ")
}
