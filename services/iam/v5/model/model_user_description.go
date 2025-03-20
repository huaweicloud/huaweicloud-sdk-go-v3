package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserDescription IAM用户描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
type UserDescription struct {
}

func (o UserDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDescription struct{}"
	}

	return strings.Join([]string{"UserDescription", string(data)}, " ")
}
