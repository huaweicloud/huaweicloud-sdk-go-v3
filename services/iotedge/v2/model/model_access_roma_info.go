package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessRomaInfo struct {

	// 认证key，加密存储
	AppKey *string `json:"app_key,omitempty" xml:"app_key"`

	// 认证secret，加密存储
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret"`
}

func (o AccessRomaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessRomaInfo struct{}"
	}

	return strings.Join([]string{"AccessRomaInfo", string(data)}, " ")
}
