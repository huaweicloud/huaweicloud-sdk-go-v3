package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 带宽信息id引用对象
type BandwidthRef struct {
	// 共享带宽的id

	Id string `json:"id"`
}

func (o BandwidthRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthRef struct{}"
	}

	return strings.Join([]string{"BandwidthRef", string(data)}, " ")
}
