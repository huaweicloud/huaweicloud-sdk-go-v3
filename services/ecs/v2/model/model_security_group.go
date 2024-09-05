package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroup 云服务器所属安全组信息。
type SecurityGroup struct {

	// 安全组名称。
	Name *string `json:"name,omitempty"`

	// 安全组ID。
	Id *string `json:"id,omitempty"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}
