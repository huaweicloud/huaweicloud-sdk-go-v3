package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// volume信息。
type Volume struct {

	// 磁盘大小。单位：GB。
	Size string `json:"size" xml:"size"`

	// 磁盘使用量。单位：GB。
	Used string `json:"used" xml:"used"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
