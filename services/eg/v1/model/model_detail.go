package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Detail 订阅的自定义事件目标参数列表
type Detail struct {

	// 自定义目标url
	Url *string `json:"url,omitempty"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	InvocationHttpParameters *InvocationHttpParameters `json:"invocation_http_parameters,omitempty"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}
