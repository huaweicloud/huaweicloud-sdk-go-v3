package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsVersion2 操作系统版本
type OsVersion2 struct {
}

func (o OsVersion2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersion2 struct{}"
	}

	return strings.Join([]string{"OsVersion2", string(data)}, " ")
}
