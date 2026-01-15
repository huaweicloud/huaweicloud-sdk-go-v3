package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityGroupIdInfo struct {

	// 安全组ID。
	Id string `json:"id"`
}

func (o SecurityGroupIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupIdInfo struct{}"
	}

	return strings.Join([]string{"SecurityGroupIdInfo", string(data)}, " ")
}
