package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodToOnDemandInstantlyReq struct {

	// |参数名称：资源ID列表。| |参数约束以及描述：该参数必填，数组范围限制:1-10，字符长度限制1-64。|
	ResourceIds []string `json:"resource_ids"`
}

func (o PeriodToOnDemandInstantlyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodToOnDemandInstantlyReq struct{}"
	}

	return strings.Join([]string{"PeriodToOnDemandInstantlyReq", string(data)}, " ")
}
