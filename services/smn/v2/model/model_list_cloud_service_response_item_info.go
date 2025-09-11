package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCloudServiceResponseItemInfo struct {

	// 云服务名称。
	Name string `json:"name"`

	// 云服务显示名称。
	ShowName string `json:"show_name"`
}

func (o ListCloudServiceResponseItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudServiceResponseItemInfo struct{}"
	}

	return strings.Join([]string{"ListCloudServiceResponseItemInfo", string(data)}, " ")
}
