package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerIpsListVo struct {

	// 动作（0：只记录日志，1：重置/拦截）
	Action *int32 `json:"action,omitempty"`

	// 操作系统
	AffectedOs *int32 `json:"affected_os,omitempty"`

	// 攻击类型
	AttackType *int32 `json:"attack_type,omitempty"`

	// 规则状态（0：初始化，1：配置中，2：配置成功，3：配置失败）
	ConfigStatus *int32 `json:"config_status,omitempty"`

	// 内容json存储
	Content *string `json:"content,omitempty"`

	// 目的端口类型
	DstPortType *int32 `json:"dst_port_type,omitempty"`

	// 目的端口
	DstPorts *string `json:"dst_ports,omitempty"`

	// 防火墙集群id
	GroupId *string `json:"group_id,omitempty"`

	// cfw侧自定义ips规则id
	IpsCfwId *string `json:"ips_cfw_id,omitempty"`

	// 山石侧规则id
	IpsId *string `json:"ips_id,omitempty"`

	// ips规则名称
	IpsName *string `json:"ips_name,omitempty"`

	// 协议
	Protocol *int32 `json:"protocol,omitempty"`

	// 严重程度（critical：致命，high：高危，medium:中危，low:低危）
	Severity *int32 `json:"severity,omitempty"`

	// 影响软件
	Software *int32 `json:"software,omitempty"`

	// 源端口类型
	SrcPortType *int32 `json:"src_port_type,omitempty"`

	// 源端口
	SrcPorts *string `json:"src_ports,omitempty"`
}

func (o CustomerIpsListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerIpsListVo struct{}"
	}

	return strings.Join([]string{"CustomerIpsListVo", string(data)}, " ")
}
