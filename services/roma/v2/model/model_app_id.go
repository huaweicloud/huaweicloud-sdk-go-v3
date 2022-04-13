package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用ID
type AppId struct {
}

func (o AppId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppId struct{}"
	}

	return strings.Join([]string{"AppId", string(data)}, " ")
}
