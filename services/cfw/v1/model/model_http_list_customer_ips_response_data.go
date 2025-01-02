package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpListCustomerIpsResponseData struct {

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 查询获得自定义ips规则列表总数
	Total *int32 `json:"total,omitempty"`

	// 自定义ips规则记录
	Records *[]CustomerIpsListVo `json:"records,omitempty"`
}

func (o HttpListCustomerIpsResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpListCustomerIpsResponseData struct{}"
	}

	return strings.Join([]string{"HttpListCustomerIpsResponseData", string(data)}, " ")
}
