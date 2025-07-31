package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsVersion1 系统版本
type OsVersion1 struct {
}

func (o OsVersion1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersion1 struct{}"
	}

	return strings.Join([]string{"OsVersion1", string(data)}, " ")
}
