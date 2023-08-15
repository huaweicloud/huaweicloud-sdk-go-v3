package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRespItem 实例信息。
type CreateInstanceRespItem struct {

	// 实例id
	Id string `json:"id"`

	// 实例名称。 用于表示实例的名称，同一租户下，同类型的实例名可重名。 取值范围：4~64个字符之间，必须以字母开头，区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符。
	Name string `json:"name"`

	// 实例状态。如BUILD，表示创建中。 仅创建按需实例时会返回该参数。
	Status string `json:"status"`

	Datastore *OpenGaussDatastoreResponse `json:"datastore"`

	Ha *OpenGaussHaResponse `json:"ha"`

	// 数据库端口信息。 当不传该参数时，默认端口8000。
	Port string `json:"port"`

	Volume *OpenGaussVolumeResponse `json:"volume"`

	// 实例副本数。
	ReplicaNum int32 `json:"replica_num"`

	// 区域ID。创建主实例时必选，其它场景不可选。 取值参见[地区和终端节点](https://developer.huaweicloud.com/endpoint)。
	Region string `json:"region"`

	BackupStrategy *OpenGaussBackupStrategyForResponse `json:"backup_strategy"`

	// 规格码。
	FlavorRef string `json:"flavor_ref"`

	// 可用区ID。GaussDB取值范围：非空，可选部署在同一可用区或三个不同可用区，可用区之间用逗号隔开。 取值参见[地区和终端节点](https://developer.huaweicloud.com/endpoint)。
	AvailabilityZone string `json:"availability_zone"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 子网ID。
	SubnetId string `json:"subnet_id"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	ChargeInfo *OpenGaussChargeInfo `json:"charge_info"`
}

func (o CreateInstanceRespItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRespItem struct{}"
	}

	return strings.Join([]string{"CreateInstanceRespItem", string(data)}, " ")
}
