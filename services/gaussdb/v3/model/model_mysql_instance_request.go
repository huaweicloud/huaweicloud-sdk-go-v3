package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例信息
type MysqlInstanceRequest struct {
	ChargeInfo *MysqlChargeInfo `json:"charge_info,omitempty" xml:"charge_info"`

	// 区域ID。
	Region string `json:"region" xml:"region"`

	// 实例名称。 用于表示实例的名称，同一租户下，同类型的实例名可重名。 取值范围：4~64个字符之间，必须以字母开头，区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符。
	Name string `json:"name" xml:"name"`

	Datastore *MysqlDatastore `json:"datastore" xml:"datastore"`

	// 实例类型，目前仅支持Cluster。
	Mode string `json:"mode" xml:"mode"`

	// 规格码。
	FlavorRef string `json:"flavor_ref" xml:"flavor_ref"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id" xml:"vpc_id"`

	// 子网的网络ID。
	SubnetId string `json:"subnet_id" xml:"subnet_id"`

	// 安全组ID。如果实例所选用的子网开启网络ACL进行访问控制，则该参数非必选。如果未开启ACL进行访问控制，则该参数必选。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 参数模板ID。
	ConfigurationId *string `json:"configuration_id,omitempty" xml:"configuration_id"`

	// 数据库密码。 取值范围：至少包含以下字符的三种：大小写字母、数字和特殊符号~!@#$%^*-_=+?,()&，长度8~32个字符。 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。如果您输入弱密码，系统会自动判定密码非法。
	Password string `json:"password" xml:"password"`

	BackupStrategy *MysqlBackupStrategy `json:"backup_strategy,omitempty" xml:"backup_strategy"`

	// 时区。默认时区为UTC。
	TimeZone *string `json:"time_zone,omitempty" xml:"time_zone"`

	// 可用区类型,单可用区single或多可用区multi。
	AvailabilityZoneMode string `json:"availability_zone_mode" xml:"availability_zone_mode"`

	// 主可用区。
	MasterAvailabilityZone *string `json:"master_availability_zone,omitempty" xml:"master_availability_zone"`

	// 备节点个数。单次接口调用最多支持创建9个备节点。
	SlaveCount int32 `json:"slave_count" xml:"slave_count"`

	Volume *MysqlVolume `json:"volume,omitempty" xml:"volume"`

	Tags *[]MysqlTags `json:"tags,omitempty" xml:"tags"`

	// 企业项目ID。如果账户开通企业项目服务则该参数必选，未开启该参数不可选。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 专属资源池ID，只有开通专属资源池后才可以下发此参数。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty" xml:"dedicated_resource_id"`
}

func (o MysqlInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlInstanceRequest struct{}"
	}

	return strings.Join([]string{"MysqlInstanceRequest", string(data)}, " ")
}
