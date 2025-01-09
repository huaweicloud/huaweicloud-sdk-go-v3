package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroup 云应用使用的安全组，如果不指定则默认使用云服务代理中指定的安全组。
type SecurityGroup struct {

	// 安全组ID。
	Id *string `json:"id,omitempty"`

	// 安全组名称。
	Name *string `json:"name,omitempty"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}
