package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceResponse Response Object
type DeleteInstanceResponse struct {

	// 独享引擎实例ID
	Id *string `json:"id,omitempty"`

	// 独享引擎实例名称
	Instancename *string `json:"instancename,omitempty"`

	// 独享引擎实例Region ID
	Region *string `json:"region,omitempty"`

	// 可用区ID
	Zone *string `json:"zone,omitempty"`

	// CPU架构
	Arch *string `json:"arch,omitempty"`

	// ECS规格
	CpuFlavor *string `json:"cpu_flavor,omitempty"`

	// 独享引擎实例所在VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 独享引擎实例所在VPC的子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 独享引擎实例的业务面IP
	ServiceIp *string `json:"service_ip,omitempty"`

	// 独享引擎绑定的安全组
	SecurityGroupIds *[]string `json:"security_group_ids,omitempty"`

	// 独享引擎计费状态   - 0：正常计费   - 1：冻结,资源和数据会保留，但租户无法再正常使用云服务   - 2：终止，资源和数据将清除
	Status *int32 `json:"status,omitempty"`

	// 独享引擎运行状态   - 0：创建中   - 1：运行中   - 2：删除中   - 3：已删除   - 4：创建失败   - 5：已冻结   - 6：异常   - 7：更新中   - 8：更新失败
	RunStatus *int32 `json:"run_status,omitempty"`

	// 独享引擎接入状态（0：未接入，1：已接入）
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 独享引擎是否可升级（0：不可升级，1：可升级）
	Upgradable *int32 `json:"upgradable,omitempty"`

	// 云服务代码。 仅作为标记，用户可忽略。
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// 云服务资源类型，仅作为标记，用户可忽略。
	ResourceType *string `json:"resourceType,omitempty"`

	// 云服务资源代码。仅作为标记，用户可忽略。
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	// 独享引擎ECS规格，如\"8vCPUs | 16GB\"
	Specification *string `json:"specification,omitempty"`

	// 独享引擎ECS ID
	ServerId *string `json:"serverId,omitempty"`

	// 引擎实例创建时间
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceResponse", string(data)}, " ")
}
