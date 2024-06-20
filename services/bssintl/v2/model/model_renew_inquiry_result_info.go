package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RenewInquiryResultInfo struct {

	// |参数名称：资源ID。| |参数约束及描述：资源ID。请求携带“include_relative_resources”字段，会返回关联资源的主资源ID和续订金额|
	ResourceId *string `json:"resource_id,omitempty"`

	// |参数名称：主资源（包含从资源）续订金额。单位为元| |参数约束及描述：主资源（包含从资源）续订金额。单位为元|
	Amount *string `json:"amount,omitempty"`
}

func (o RenewInquiryResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewInquiryResultInfo struct{}"
	}

	return strings.Join([]string{"RenewInquiryResultInfo", string(data)}, " ")
}
