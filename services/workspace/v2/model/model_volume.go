package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Volume 磁盘。
type Volume struct {

	// 磁盘类型。
	Type string `json:"type"`

	// 磁盘大小，单位GB。
	Size int64 `json:"size"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
