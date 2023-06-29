package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsVersionInfo 操作系统详情
type OsVersionInfo struct {

	// 操作系统的平台值
	Platform string `json:"platform"`

	// os_version的key值，和os_version值相同
	OsVersionKey string `json:"os_version_key"`

	// 操作系统的版本
	OsVersion string `json:"os_version"`

	// 操作系统的位数
	OsBit int32 `json:"os_bit"`

	// 操作系统的类型，Linux或Windows
	OsType string `json:"os_type"`
}

func (o OsVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersionInfo struct{}"
	}

	return strings.Join([]string{"OsVersionInfo", string(data)}, " ")
}
