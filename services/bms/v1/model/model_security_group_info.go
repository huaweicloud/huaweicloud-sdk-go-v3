package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroupInfo 安全组信息
type SecurityGroupInfo struct {

	// 安全组ID，UUID格式。
	Id string `json:"id"`
}

func (o SecurityGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupInfo struct{}"
	}

	return strings.Join([]string{"SecurityGroupInfo", string(data)}, " ")
}
