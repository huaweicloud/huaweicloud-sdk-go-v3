package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapUserName Iam用户名称
type IsapUserName struct {
}

func (o IsapUserName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapUserName struct{}"
	}

	return strings.Join([]string{"IsapUserName", string(data)}, " ")
}
