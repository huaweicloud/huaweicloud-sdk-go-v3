package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityGroupInfo struct {

	// 安全组ID。
	Id *string `json:"id,omitempty"`

	// 安全组名称。
	Name *string `json:"name,omitempty"`
}

func (o SecurityGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupInfo struct{}"
	}

	return strings.Join([]string{"SecurityGroupInfo", string(data)}, " ")
}
