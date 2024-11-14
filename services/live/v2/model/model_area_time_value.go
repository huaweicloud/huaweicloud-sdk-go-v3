package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AreaTimeValue struct {

	// 各个大区下的具体省份、区域、国家的名称。  中国内地返回结果为省份/直辖市的中文名称，比如：广东、上海； 海外大区下的地区/国家对应关系请参考[地区/国家代码对照表](https://support.huaweicloud.com/api-live/live_03_0049.html)。
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
