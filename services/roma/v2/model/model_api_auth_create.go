package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiAuthCreate struct {

	// 需要授权的环境编号
	EnvId string `json:"env_id"`

	// APP的编号列表
	AppIds []string `json:"app_ids"`

	// API的编号列表。
	ApiIds []string `json:"api_ids"`

	// 授权通道类型： - GREEN：绿色通道 - NORMAL：非绿色通道  实例开启green_tunnel特性时可以开启绿色通道，此字段不填默认为不使用绿色通道
	AuthTunnel *string `json:"auth_tunnel,omitempty"`

	// 绿色通道授权白名单。  允许白名单中的IP不使用认证信息访问，auth_tunnel = GREEN时生效
	AuthWhitelist *[]string `json:"auth_whitelist,omitempty"`

	// 绿色通道授权黑名单。  auth_tunnel = GREEN时生效
	AuthBlacklist *[]string `json:"auth_blacklist,omitempty"`

	// 访问参数列表。
	VisitParams *[]ApiAuthVisitParam `json:"visit_params,omitempty"`
}

func (o ApiAuthCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAuthCreate struct{}"
	}

	return strings.Join([]string{"ApiAuthCreate", string(data)}, " ")
}
