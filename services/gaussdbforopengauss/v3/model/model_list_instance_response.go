package model

import (
	"encoding/json"

	"strings"
)

// 实例信息。
type ListInstanceResponse struct {
	// 实例ID。

	Id string `json:"id"`
	// 创建的实例名称。

	Name string `json:"name"`
	// 实例状态。  取值：  值为“BUILD”，表示实例正在创建。  值为“ACTIVE”，表示实例正常。  值为“FAILED”，表示实例异常。  值为“FROZEN”，表示实例冻结。  值为“EXPANDING”，表示实例正在扩容CN或DN。  值为“REBOOTING”，表示实例正在重启。  值为“UPGRADING”，表示实例正在升级。  值为“BACKING UP”，表示实例正在进行备份。  值为“STORAGE FULL”，表示实例磁盘空间满。

	Status string `json:"status"`
	// 实例内网IP地址列表。CN在的弹性云服务器创建成功后该值存在，其他情况下为空字符串。

	PrivateIps []string `json:"private_ips"`
	// 实例外网IP地址列表。绑定弹性公网IP后，该值不为空。

	PublicIps []string `json:"public_ips"`
	// 数据库端口号。GaussDB(for openGauss)数据库端口设置范围为1024~39998（其中2378,2379,2380,4999,5000,5999,6000,6001,8097,8098,20049,20050,21731,21732被系统占用不可设置）。  当不传该参数时，默认端口如下：8000。

	Port int32 `json:"port"`
	// 实例类型，取值为 \"Enterprise\"，对应于分布式实例（企业版）。

	Type string `json:"type"`

	Ha *ListHa `json:"ha"`
	// 实例副本数。说明：需添加白名单，才会开放显示。

	ReplicaNum *int32 `json:"replica_num,omitempty"`
	// 实例所在区域。

	Region string `json:"region"`

	Datastore *ListDatastore `json:"datastore"`
	// 创建时间，格式为“yyyy-mm-dd hh:mm:ss timezone”。  其中timezone是指时区。  说明：创建时该值为实例下发创建的时间，创建完成后，该值为创建完成时间。

	Created string `json:"created"`
	// 更新时间，格式与“created”字段对应格式完全相同。  说明：创建时返回值为空，数据库实例创建成功后该值不为空。

	Updated string `json:"updated"`
	// 默认用户名。

	DbUserName string `json:"db_user_name"`
	// 虚拟私有云ID。

	VpcId string `json:"vpc_id"`
	// 子网的网络ID信息。

	SubnetId string `json:"subnet_id"`
	// 安全组ID。

	SecurityGroupId string `json:"security_group_id"`
	// 规格码。参考[表1](https://support.huaweicloud.com/api-opengauss/opengauss_api_0037.html#opengauss_api_0037__ted9b9d433c8a4c52884e199e17f94479)中GaussDB(for openGauss)的“规格编码”列内容获取。

	FlavorRef string `json:"flavor_ref"`

	FlavorInfo *ListFlavorInfo `json:"flavor_info"`

	Volume *ListVolume `json:"volume"`
	// 数据库切换策略。取值为“reliability”或“availability”，分别对应于可靠性优先和可用性优先。

	SwitchStrategy string `json:"switch_strategy"`

	BackupStrategy *OpenGaussBackupStrategyForListResponse `json:"backup_strategy"`
	// 可维护时间窗，为UTC时间。

	MaintenanceWindow string `json:"maintenance_window"`
	// 所关联的数据库实例列表。GaussDB(for openGauss)不涉及该参数。

	RelatedInstance []interface{} `json:"related_instance"`
	// 实例节点信息。

	Nodes []interface{} `json:"nodes"`
	// 企业项目标签ID。非企业项目账号的实例，企业项目默认0。

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// 磁盘加密密钥ID。

	DiskEncryptionId string `json:"disk_encryption_id"`

	ChargeInfo *OpenGaussChargeInfoListResponse `json:"charge_info"`
	// 时区。

	TimeZone string `json:"time_zone"`
	// 标签列表，没有标签默认为空数组。

	Tags []interface{} `json:"tags"`
}

func (o ListInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceResponse", string(data)}, " ")
}
