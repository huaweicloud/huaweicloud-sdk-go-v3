package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenGaussInstanceResult 实例信息。
type OpenGaussInstanceResult struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。用于表示实例的名称，同一租户下，同类型的实例名称可相同。  取值范围：4~64个字符之间，必须以字母开头，不区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符。
	Name string `json:"name"`

	// 实例状态。如BUILD，表示创建中。  仅创建按需实例时会返回该参数。
	Status string `json:"status"`

	Datastore *OpenGaussDatastoreResult `json:"datastore"`

	Ha *OpenGaussHaResult `json:"ha,omitempty"`

	// 实例副本数。
	ReplicaNum *int32 `json:"replica_num,omitempty"`

	BackupStrategy *OpenGaussBackupStrategyForResponse `json:"backup_strategy"`

	// 数据库端口信息，与请求参数相同。
	Port string `json:"port"`

	// 项目标签。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 规格码，取值范围：非空。参考[表1](https://support.huaweicloud.com/api-opengauss/opengauss_api_0037.html#opengauss_api_0037__ted9b9d433c8a4c52884e199e17f94479)中GaussDB的“规格编码”列内容获取。
	FlavorRef string `json:"flavor_ref"`

	Volume *OpenGaussVolumeResult `json:"volume"`

	// 区域ID。
	Region string `json:"region"`

	// 可用区ID。
	AvailabilityZone string `json:"availability_zone"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 子网的网络ID信息。
	SubnetId string `json:"subnet_id"`

	// 实例所属的安全组。
	SecurityGroupId string `json:"security_group_id"`

	ChargeInfo *OpenGaussChargeInfoResponse `json:"charge_info"`
}

func (o OpenGaussInstanceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussInstanceResult struct{}"
	}

	return strings.Join([]string{"OpenGaussInstanceResult", string(data)}, " ")
}
