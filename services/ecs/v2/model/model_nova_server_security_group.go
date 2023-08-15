package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaServerSecurityGroup
type NovaServerSecurityGroup struct {

	// 安全组名称或者uuid。
	Name *string `json:"name,omitempty"`
}

func (o NovaServerSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerSecurityGroup struct{}"
	}

	return strings.Join([]string{"NovaServerSecurityGroup", string(data)}, " ")
}
