package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AreaTimeValue struct {
	// 各个大区下的具体省份、区域、国家的名称。  具体取值请参考[国家名称缩写](vod_08_0172.xml)和[省份名称缩写](live_03_0043.xml)。

	Name string `json:"name"`
	// 当前时间返回指定指标的值

	Data []TimeValue `json:"data"`
}

func (o AreaTimeValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AreaTimeValue struct{}"
	}

	return strings.Join([]string{"AreaTimeValue", string(data)}, " ")
}
