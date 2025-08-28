package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupDescription 用户组描述信息，长度为0到255字符，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
type GroupDescription struct {
}

func (o GroupDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupDescription struct{}"
	}

	return strings.Join([]string{"GroupDescription", string(data)}, " ")
}
