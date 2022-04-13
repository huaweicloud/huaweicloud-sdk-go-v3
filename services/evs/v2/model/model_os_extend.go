package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type OsExtend struct {
	// 扩容后的云硬盘大小，单位为GB。扩容的大小必须大于原有云硬盘容量且小于云硬盘最大容量。 云硬盘最大容量： * 数据盘：32768GB * 系统盘：1024GB

	NewSize int32 `json:"new_size"`
}

func (o OsExtend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsExtend struct{}"
	}

	return strings.Join([]string{"OsExtend", string(data)}, " ")
}
