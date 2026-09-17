package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessRomaBriefInfo ROMA接入鉴权信息
type AccessRomaBriefInfo struct {

	// 认证key，加密存储
	AppKey *string `json:"app_key,omitempty"`
}

func (o AccessRomaBriefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessRomaBriefInfo struct{}"
	}

	return strings.Join([]string{"AccessRomaBriefInfo", string(data)}, " ")
}
