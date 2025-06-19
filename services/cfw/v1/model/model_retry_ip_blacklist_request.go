package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryIpBlacklistRequest Request Object
type RetryIpBlacklistRequest struct {

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	// 指定导入失败的IP黑名单的名字
	Name *string `json:"name,omitempty"`
}

func (o RetryIpBlacklistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryIpBlacklistRequest struct{}"
	}

	return strings.Join([]string{"RetryIpBlacklistRequest", string(data)}, " ")
}
