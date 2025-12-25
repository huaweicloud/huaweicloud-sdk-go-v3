package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapScriptName 脚本名称
type IsapScriptName struct {
}

func (o IsapScriptName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapScriptName struct{}"
	}

	return strings.Join([]string{"IsapScriptName", string(data)}, " ")
}
