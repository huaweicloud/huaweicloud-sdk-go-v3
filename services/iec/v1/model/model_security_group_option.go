package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type SecurityGroupOption struct {
	// 边缘实例的安全组，会对边缘实例中配置的网卡生效。需要指定已有安全组的ID。不填写时选择默认安全组

	Id *string `json:"id,omitempty"`
}

func (o SecurityGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupOption struct{}"
	}

	return strings.Join([]string{"SecurityGroupOption", string(data)}, " ")
}
