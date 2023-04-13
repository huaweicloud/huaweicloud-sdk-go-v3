package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnoseResourceVo struct {

	// ID
	Fault *string `json:"fault,omitempty"`

	// 状态
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// ip4地址
	AccessIPv4 *string `json:"accessIPv4,omitempty"`

	// ip6地址
	AccessIPv6 *string `json:"accessIPv6,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 进度
	Progress *string `json:"progress,omitempty"`

	// 前端id
	TenantId *string `json:"tenantId,omitempty"`

	// 用户id
	UserId *string `json:"userId,omitempty"`

	// 资源元数据
	Metadata *interface{} `json:"metadata,omitempty"`

	// 主机id
	HostId *string `json:"hostId,omitempty"`
}

func (o DiagnoseResourceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnoseResourceVo struct{}"
	}

	return strings.Join([]string{"DiagnoseResourceVo", string(data)}, " ")
}
