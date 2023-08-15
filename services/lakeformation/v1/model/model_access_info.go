package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessInfo 接入信息
type AccessInfo struct {

	// 服务名称
	VpcepServiceName *string `json:"vpcep_service_name,omitempty"`

	// 分组独立域名
	Domain *string `json:"domain,omitempty"`
}

func (o AccessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessInfo struct{}"
	}

	return strings.Join([]string{"AccessInfo", string(data)}, " ")
}
