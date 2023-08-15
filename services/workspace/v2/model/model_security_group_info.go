package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityGroupInfo struct {

	// 安全组ID。
	Id string `json:"id"`
}

func (o SecurityGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupInfo struct{}"
	}

	return strings.Join([]string{"SecurityGroupInfo", string(data)}, " ")
}
