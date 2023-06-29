package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessRequestInfo 申请接入服务的请求信息
type AccessRequestInfo struct {

	// 服务创建的id
	VpcepId string `json:"vpcep_id"`

	// 分组独立域名
	Domain string `json:"domain"`
}

func (o AccessRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessRequestInfo struct{}"
	}

	return strings.Join([]string{"AccessRequestInfo", string(data)}, " ")
}
