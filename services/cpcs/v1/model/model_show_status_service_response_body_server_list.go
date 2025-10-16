package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStatusServiceResponseBodyServerList struct {

	// 服务类型
	ServerType *string `json:"server_type,omitempty"`

	// 应用列表
	AppList *[]ShowStatusAppItem `json:"app_list,omitempty"`
}

func (o ShowStatusServiceResponseBodyServerList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusServiceResponseBodyServerList struct{}"
	}

	return strings.Join([]string{"ShowStatusServiceResponseBodyServerList", string(data)}, " ")
}
