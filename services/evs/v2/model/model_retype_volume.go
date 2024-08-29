package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetypeVolume 变更磁盘类型
type RetypeVolume struct {

	// 磁盘变更至指定的磁盘类型
	NewType string `json:"new_type"`
}

func (o RetypeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetypeVolume struct{}"
	}

	return strings.Join([]string{"RetypeVolume", string(data)}, " ")
}
