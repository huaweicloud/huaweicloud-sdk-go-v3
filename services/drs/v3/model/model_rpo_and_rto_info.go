package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RpoAndRTO信息体
type RpoAndRtoInfo struct {

	// 检查点
	CheckPoint *string `json:"check_point,omitempty" xml:"check_point"`

	// 延迟
	Delay *string `json:"delay,omitempty" xml:"delay"`

	// gtid集合
	GtidSet *string `json:"gtid_set,omitempty" xml:"gtid_set"`

	// 当前时间 ，格式为“yyyy-MM-dd HH:mm:ss”
	Time *string `json:"time,omitempty" xml:"time"`
}

func (o RpoAndRtoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RpoAndRtoInfo struct{}"
	}

	return strings.Join([]string{"RpoAndRtoInfo", string(data)}, " ")
}
