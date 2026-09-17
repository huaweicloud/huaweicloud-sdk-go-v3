package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LicenseInfo license信息
type LicenseInfo struct {

	// esn码
	Esn string `json:"esn"`

	// 超期时间
	ExpireTime string `json:"expire_time"`

	// 宽限期
	GraceTime *string `json:"grace_time,omitempty"`
}

func (o LicenseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LicenseInfo struct{}"
	}

	return strings.Join([]string{"LicenseInfo", string(data)}, " ")
}
