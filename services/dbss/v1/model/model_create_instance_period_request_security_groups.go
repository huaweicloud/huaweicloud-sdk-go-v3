package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstancePeriodRequestSecurityGroups struct {

	// 云服务器对应的安全组ID，会对创建云服务器中配置的网卡生效
	Id string `json:"id"`
}

func (o CreateInstancePeriodRequestSecurityGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancePeriodRequestSecurityGroups struct{}"
	}

	return strings.Join([]string{"CreateInstancePeriodRequestSecurityGroups", string(data)}, " ")
}
