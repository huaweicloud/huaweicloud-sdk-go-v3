package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsVersion 系统版本
type OsVersion struct {
}

func (o OsVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersion struct{}"
	}

	return strings.Join([]string{"OsVersion", string(data)}, " ")
}
