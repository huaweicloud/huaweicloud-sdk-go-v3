package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWtpProtectHostRequest Request Object
type ListWtpProtectHostRequest struct {

	// Region Id
	Region string `json:"region"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 主机ID
	HostId *string `json:"host_id,omitempty"`

	// 弹性公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty"`

	// 操作系统类别（linux，windows）   - linux : linux操作系统   - windows : windows操作系统
	OsType *string `json:"os_type,omitempty"`

	// 配额状态   - opened : 已绑定网页防篡改配额
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 网页防篡改防护状态   - opened : 防护汇总   - opening : 正在开启   - open_failed : 防护失败   - partial_protection : 部分防护   - protection_interruption : 防护中断
	WtpStatus *string `json:"wtp_status,omitempty"`

	// 客户端状态   - not_installed : agent未安装   - online : agent在线   - offline : agent不在线
	AgentStatus *string `json:"agent_status,omitempty"`

	// 默认10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWtpProtectHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWtpProtectHostRequest struct{}"
	}

	return strings.Join([]string{"ListWtpProtectHostRequest", string(data)}, " ")
}
