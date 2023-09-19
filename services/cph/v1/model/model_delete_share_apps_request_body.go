package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteShareAppsRequestBody struct {

	// 所需删除的共享应用的合法包名。最大长度128字节。只支持包含大小写字母、数字、下划线、点，其中不允许以数字和下划线开头，点不能作为结尾且包名中至少有一个点。
	PackageName string `json:"package_name"`

	// 云手机服务器ID列表。
	ServerIds []string `json:"server_ids"`
}

func (o DeleteShareAppsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareAppsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteShareAppsRequestBody", string(data)}, " ")
}
