package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEdgeSiteResourcesVolume 磁盘。
type CheckEdgeSiteResourcesVolume struct {

	// 磁盘类型。
	Type string `json:"type"`

	// 磁盘大小，单位GB。
	Size int64 `json:"size"`
}

func (o CheckEdgeSiteResourcesVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEdgeSiteResourcesVolume struct{}"
	}

	return strings.Join([]string{"CheckEdgeSiteResourcesVolume", string(data)}, " ")
}
