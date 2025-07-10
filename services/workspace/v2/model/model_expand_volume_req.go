package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandVolumeReq 扩容磁盘请求。
type ExpandVolumeReq struct {

	// 扩容后的磁盘大小，单位为GB。
	NewSize int32 `json:"new_size"`
}

func (o ExpandVolumeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandVolumeReq struct{}"
	}

	return strings.Join([]string{"ExpandVolumeReq", string(data)}, " ")
}
