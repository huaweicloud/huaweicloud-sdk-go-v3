package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteShareAppsRequestBody struct {

	// 所需删除的共享应用的合法包名
	PackageName string `json:"package_name" xml:"package_name"`

	// 云手机服务器ID列表
	ServerIds []string `json:"server_ids" xml:"server_ids"`
}

func (o DeleteShareAppsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareAppsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteShareAppsRequestBody", string(data)}, " ")
}
