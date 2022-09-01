package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PartnersServiceInfo struct {

	// 客户id
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 客户名称
	CustomerName *string `json:"customer_name,omitempty" xml:"customer_name"`

	// 服务时区，GMT+08:00
	ServiceTimeZone *string `json:"service_time_zone,omitempty" xml:"service_time_zone"`

	// 每周服务天数
	ServiceTimeDay *string `json:"service_time_day,omitempty" xml:"service_time_day"`

	// 每天服务小时
	ServiceTimeHour *string `json:"service_time_hour,omitempty" xml:"service_time_hour"`
}

func (o PartnersServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartnersServiceInfo struct{}"
	}

	return strings.Join([]string{"PartnersServiceInfo", string(data)}, " ")
}
