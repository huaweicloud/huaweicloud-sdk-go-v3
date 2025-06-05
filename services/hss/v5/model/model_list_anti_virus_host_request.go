package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusHostRequest Request Object
type ListAntiVirusHostRequest struct {

	// **参数解释**: 主机所属的企业项目ID。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。 **约束限制**: 开通企业项目功能后才需要配置企业项目。 **取值范围**: 字符长度1-256位 **默认取值**: 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset int32 `json:"offset"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty"`

	// 任务类型，包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
	ScanType string `json:"scan_type"`

	// 启动类型，包含如下:   - now ：立即启动   - period : 周期启动
	StartType string `json:"start_type"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 下次启动时间，毫秒
	NextStartTime *int64 `json:"next_start_time,omitempty"`
}

func (o ListAntiVirusHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiVirusHostRequest struct{}"
	}

	return strings.Join([]string{"ListAntiVirusHostRequest", string(data)}, " ")
}
