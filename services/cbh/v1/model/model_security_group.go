package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 安全组
type SecurityGroup struct {

	// 安全组ID
	Id string `json:"id"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}
