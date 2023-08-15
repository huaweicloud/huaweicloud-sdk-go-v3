package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppInfo struct {

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`
}

func (o AppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInfo struct{}"
	}

	return strings.Join([]string{"AppInfo", string(data)}, " ")
}
