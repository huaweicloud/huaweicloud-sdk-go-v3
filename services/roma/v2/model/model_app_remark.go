package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用描述
type AppRemark struct {
}

func (o AppRemark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppRemark struct{}"
	}

	return strings.Join([]string{"AppRemark", string(data)}, " ")
}
