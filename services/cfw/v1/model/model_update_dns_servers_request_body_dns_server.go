package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDnsServersRequestBodyDnsServer struct {

	// DNS服务器IP，可通过[查询dns服务器列表接口](ListDnsServers.xml)查询获得，通过返回值中的data.server_ip（.表示各对象之间层级的区分）获得。
	ServerIp string `json:"server_ip"`

	// 是否是用户自定义的dns服务器，0否 1是
	IsCustomized int32 `json:"is_customized"`

	// 是否应用，0否 1是
	IsApplied int32 `json:"is_applied"`
}

func (o UpdateDnsServersRequestBodyDnsServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersRequestBodyDnsServer struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersRequestBodyDnsServer", string(data)}, " ")
}
