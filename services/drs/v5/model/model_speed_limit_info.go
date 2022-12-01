package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 限速信息体。 - 限速：自定义的最大迁移速度，迁移过程中的迁移速度将不会超过该速度。 - 不限速：对迁移速度不进行限制，通常会最大化使用源数据库的出口带宽。该流速模式同时会对源数据库造成读消耗，消耗取决于源数据库的出口带宽。比如：源数据库的出口带宽为100MB/s，假设高速模式使用了80%带宽，则迁移对源数据库将造成80MB/s的读操作IO消耗。
type SpeedLimitInfo struct {

	// 开始限速时间，此时间为UTC时间，开始时间为整时，若有分钟，则会忽略，格式为hh:mm，小时数为两位，例如：01:00。
	Begin string `json:"begin"`

	// 结束时间，此时间为UTC时间，输入必须为59分结尾，格式为hh:mm，例如：15:59。
	End string `json:"end"`

	// 限速值，取值范围为1~9999，单位为MB/s。
	Speed string `json:"speed"`
}

func (o SpeedLimitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeedLimitInfo struct{}"
	}

	return strings.Join([]string{"SpeedLimitInfo", string(data)}, " ")
}
