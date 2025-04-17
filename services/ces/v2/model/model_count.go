package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Count 告警连续触发次数，事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180
type Count struct {
}

func (o Count) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Count struct{}"
	}

	return strings.Join([]string{"Count", string(data)}, " ")
}
