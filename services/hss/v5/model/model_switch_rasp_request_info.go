package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchRaspRequestInfo struct {

	// HostId list
	HostIdList []string `json:"host_id_list"`

	// 应用防护开关状态
	AppStatus bool `json:"app_status"`

	// 应用防护类型
	AppType *string `json:"app_type,omitempty"`

	// 防护策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 动态防护开关状态
	AutoAttach *bool `json:"auto_attach,omitempty"`

	// RASP端口列表，列表元素与host_id_list对应，app_status为true时支持生效
	RaspPortList *[]int32 `json:"rasp_port_list,omitempty"`
}

func (o SwitchRaspRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchRaspRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchRaspRequestInfo", string(data)}, " ")
}
