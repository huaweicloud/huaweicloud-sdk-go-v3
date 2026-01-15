package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseServer struct {

	// 连接IP
	ConnectIp *string `json:"connect_ip,omitempty"`

	// CPU数
	Cpu *string `json:"cpu,omitempty"`

	// 创建时间
	Created *string `json:"created,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 实例冻结状态  - 1：冻结可释放 （默认）  - 2：冻结不可释放  - 3：冻结后不可续费
	Effect *int32 `json:"effect,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 双机实例HA中用来标注实例为主机还是备机, - 0：主机  - 1：备机
	IsActive *int32 `json:"is_active,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 内存大小
	Ram *string `json:"ram,omitempty"`

	// 所属REGION
	Region *string `json:"region,omitempty"`

	// 实例所属规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 冻结场景  - POLICE：公安冻结  - ILLEGAL：违规冻结  - VERIFY：未实名认证冻结  - PARTNER：合作伙伴冻结 - ARREAR：普通冻结（普通）
	Scene *string `json:"scene,omitempty"`

	// 安全组ID
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 资源规格类型编码
	Specification *string `json:"specification,omitempty"`

	// 实例状态
	Status *string `json:"status,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 实例处理中状态
	TaskStatus *string `json:"task_status,omitempty"`

	// 升级状态  - SUCCESS：成功  - FAILED：失败  - UPGRADING：升级中
	UpdateStatus *string `json:"update_status,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 所属区域
	Zone *string `json:"zone,omitempty"`

	// 服务器ID
	ServerId *string `json:"server_id,omitempty"`
}

func (o ResponseServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseServer struct{}"
	}

	return strings.Join([]string{"ResponseServer", string(data)}, " ")
}
