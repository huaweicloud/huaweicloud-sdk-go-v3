package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 带宽对象
type BatchCreateBandwidthOption struct {

	// 取值范围：正整数  功能说明：批创的共享带宽的个数  说明： 如果传入的参数为小数（如 2.2）或者字符类型（如“2”），会自动强制转换为整数。
	Count int32 `json:"count" xml:"count"`

	// 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）  功能说明：带宽名称
	Name string `json:"name" xml:"name"`

	// 功能说明：带宽大小。共享带宽的大小有最小值限制，默认为2M，可能因局点不同而不同。  取值范围：默认5Mbit/s~2000Mbit/s（具体范围以各区域配置为准，请参见控制台对应页面显示）。  注意：调整带宽时的最小单位会根据带宽范围不同存在差异。  小于等于300Mbit/s：默认最小单位为1Mbit/s。  300Mbit/s~1000Mbit/s：默认最小单位为50Mbit/s。  大于1000Mbit/s：默认最小单位为500Mbit/s。
	Size int32 `json:"size" xml:"size"`

	// 功能说明：表示中心站点资源或者边缘站点资源 取值范围： center、边缘站点名称 约束：共享带宽只能插入与该字段相同的publicip
	PublicBorderGroup *string `json:"public_border_group,omitempty" xml:"public_border_group"`
}

func (o BatchCreateBandwidthOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateBandwidthOption struct{}"
	}

	return strings.Join([]string{"BatchCreateBandwidthOption", string(data)}, " ")
}
